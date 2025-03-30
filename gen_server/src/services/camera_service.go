package services

import (
	"context"
	"errors"
	"fmt"
	"gen_server/src/client"
	"gen_server/src/db"
	api "gen_server/src/generated"
	"log"
	"slices"
	"strconv"

	"github.com/go-redis/redis"
)

type CameraService struct {
	queries *db.Queries
	redis   *redis.Client
}

func NewCameraService(queries *db.Queries, redis *redis.Client) *CameraService {
	return &CameraService{queries, redis}
}

func (cc *CameraService) GetCameras(ctx context.Context, params *api.ListCamerasParams) (infos *api.CameraInfos, err error) {
	idAddr, err := cc.getIdAddress(params.GymId)
	if err != nil {
		return nil, err
	}
	hostname := "http://" + idAddr + ":8080"
	log.Printf("Got hostname %s", hostname)
	localClient, err := client.NewClient(hostname)
	if err != nil {
		return nil, err
	}
	cameras, err := localClient.GetCameras(ctx)
	if err != nil {
		return nil, err
	}
	infos = &api.CameraInfos{}
	for _, item := range cameras.Cameras {
		occupiedBy, isCameraOccupied, err := cc.isCameraOccupied(ctx, int64(params.GymId), int64(item.CameraID))
		if err != nil {
			return nil, err
		}
		camera := api.CameraInfo{
			CameraID:    api.ID(item.CameraID),
			Description: item.Description,
			OccupiedBy: api.OptString{
				Value: occupiedBy,
				Set:   isCameraOccupied,
			},
		}
		infos.Cameras = append(infos.Cameras, camera)
	}
	return infos, nil
}

func (cc *CameraService) StartCameraAction(ctx context.Context, params *api.StartCameraActionParams, action *api.CameraAction) (err error) {
	// TODO: Проверка сессий
	idAddr, err := cc.getIdAddress(params.GymId)
	if err != nil {
		return err
	}
	localClient, err := client.NewClient(fmt.Sprintf("http://%s:8080/cameras", idAddr))
	if err != nil {
		return err
	}
	return localClient.StartCameraAction(
		ctx,
		&client.ActionParams{
			Velocity: client.ActionParamsVelocity{
				Pan:  action.Velocity.Pan,
				Tilt: action.Velocity.Tilt,
				Zoom: action.Velocity.Zoom,
			},
			Deadline: action.Deadline,
		},
		client.StartCameraActionParams{
			CameraId: client.ID(params.CameraId),
		},
	)
}

func (cc *CameraService) StopCameraAction(ctx context.Context, params *api.StopCameraActionParams) (err error) {
	idAddr, err := cc.getIdAddress(params.GymId)
	if err != nil {
		return err
	}
	localClient, err := client.NewClient(fmt.Sprintf("http://%s:8080/cameras", idAddr))
	if err != nil {
		return err
	}
	return localClient.StopCameraAction(ctx, client.StopCameraActionParams{
		CameraId: client.ID(params.CameraId),
	})
}

func (cc *CameraService) isCameraOccupied(ctx context.Context, gymId int64, cameraId int64) (occupiedBy string, isOccupied bool, err error) {
	occupiedCams, err := cc.queries.SelectOccupiedCams(ctx, gymId)
	if err != nil {
		return "", false, err
	}
	occupationId := slices.IndexFunc(occupiedCams, func(c db.SelectOccupiedCamsRow) bool {
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

func (cc *CameraService) getIdAddress(gymId api.ID) (string, error) {
	addr := cc.redis.Get(strconv.FormatInt(int64(gymId), 10))
	if addr == nil {
		return "", errors.New("ip address was not found in reddis")
	}
	return addr.Val(), nil
}
