package routes

import (
	"github.com/gin-gonic/gin"
	"server/src/controllers"
)

type GymRoute struct {
	gymController controllers.GymController
}

func NewRoute(gymController controllers.GymController) GymRoute {
	return GymRoute{gymController}
}

func (r *GymRoute) Route(rg *gin.RouterGroup) {
	router := rg.Group("gyms")
	router.POST("/create", r.gymController.CreateGym)
	router.GET("/get/:gym_id", r.gymController.GetGym)
}
