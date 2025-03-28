package routes

import (
	"github.com/gin-gonic/gin"
	"server/src/controllers"
)

type GymCameraRoute struct {
	gymCameraController controllers.GymCameraController
}

func NewGymCameraRoute(gymCameraController controllers.GymCameraController) GymCameraRoute {
	return GymCameraRoute{gymCameraController}
}

func (r *GymCameraRoute) Route(rg *gin.RouterGroup) {
	router := rg.Group("gym/camera")
	router.GET("/:gym_id/list", r.gymCameraController.GetCameras)
	router.POST("/ptz/:gym_id/:camera_id", r.gymCameraController.PostPtz)
	router.DELETE("/ptz/:gym_id/:camera_id", r.gymCameraController.DeletePtz)
}
