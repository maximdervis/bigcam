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
	// TODO: ptz
	// TODO: [crit] Добавить проксирование по адресу из pg
	router.GET("/list", r.gymCameraController.GetCameras)
}
