package controllers

import (
	"context"
	"errors"
	"net/http"
	"server/src/db"
	"server/src/util"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
)

type GymCameraController struct {
	db    *db.Queries
	redis *redis.Client
	ctx   context.Context
}

func NewGymCameraController(db *db.Queries, redis *redis.Client, ctx context.Context) *GymCameraController {
	return &GymCameraController{db, redis, ctx}
}

func (cc *GymCameraController) GetCameras(ctx *gin.Context) {
	// TODO: Доступы на получение камер в зале
	gymIdStr := ctx.Param("gym_id")
	gymIdInt, err := strconv.Atoi(gymIdStr)
	if err != nil {
		util.SetBadRequestStatus(ctx, err)
	}

	client := resty.New()
	var respBody map[string]interface{}
	_, err = client.R().
		SetResult(&respBody).
		Get(makeLocalUrl(cc, ctx, int64(gymIdInt), "cameras"))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}

	camerasResponse := []gin.H{}
	for _, item := range respBody {
		if cameraData, ok := item.(map[string]interface{}); ok {
			cameraId := cameraData["id"].(int64)
			occupiedBy, isCameraOccupied, err := isCameraOccupied(cc, ctx, int64(gymIdInt), cameraId) // TODO: Убрать каст в int64
			if err != nil {
				util.SetInternalErrorStatus(ctx, err)
				return
			}
			var maybeOccupiedBy *string
			if isCameraOccupied {
				maybeOccupiedBy = &occupiedBy
			}
			camera := gin.H{
				"id":          cameraData["id"],
				"description": cameraData["description"],
				"occupied_by": maybeOccupiedBy,
			}
			camerasResponse = append(camerasResponse, camera)
		} else {
			util.SetInternalErrorStatus(ctx, "unexpected data format")
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cameras": camerasResponse,
	})
}

func (cc *GymCameraController) PostPtz(ctx *gin.Context) {
	// TODO: Проверка сессий
	var err error
	cameraIdStr := ctx.Param("camera_id")
	gymIdStr := ctx.Param("gym_id")
	gymIdInt, err := strconv.Atoi(gymIdStr)
	if err != nil {
		util.SetBadRequestStatus(ctx, err)
	}
	type Velocity struct {
	  Pan  float64 `json:"pan"`
	  Tilt float64 `json:"tilt"`
	  Zoom float64 `json:"zoom"`
	}

	type Request struct {
	  Velocity Velocity `json:"velocity"`
	  Deadline string   `json:"deadline"`
	}
	var payload *Request
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		util.SetBadRequestStatus(ctx, err)
		return
	}

	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(makeLocalUrl(cc, ctx, int64(gymIdInt), "cameras/" + cameraIdStr + "/ptz"))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (cc *GymCameraController) DeletePtz(ctx *gin.Context) {
	// TODO: Проверка сессий
	var err error
	cameraIdStr := ctx.Param("camera_id")
	gymIdStr := ctx.Param("gym_id")
	gymIdInt, err := strconv.Atoi(gymIdStr)
	if err != nil {
		util.SetBadRequestStatus(ctx, err)
	}

	client := resty.New()
	_, err = client.R().
		Delete(makeLocalUrl(cc, ctx, int64(gymIdInt), "cameras/" + cameraIdStr + "/ptz"))
	if err != nil {
		util.SetInternalErrorStatus(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func isCameraOccupied(cc *GymCameraController, ctx *gin.Context, gymId int64, cameraId int64) (string, bool, error) {
			occupiedCams, err := cc.db.SelectOccupiedCams(ctx, gymId)
			if err != nil {
				return "", false, err
			}
			occupationId := slices.IndexFunc(occupiedCams, func (c db.SelectOccupiedCamsRow) (bool) {
								return c.CameraID == cameraId
			})
			if occupationId != -1 {
				occupiedCam := occupiedCams[occupationId]
				if !occupiedCam.Name.Valid {
					return "", false, errors.New("failed to find user for session")
				}
				return occupiedCams[occupationId].Name.String, true, nil
			}
			return "", false, nil
}

func makeLocalUrl(cc *GymCameraController, ctx *gin.Context, gymId int64, path string) (string) {
	ip_addr := cc.redis.Get(ctx, strconv.FormatInt(gymId, 10))
	return "http://"+ip_addr.Val()+path
}
