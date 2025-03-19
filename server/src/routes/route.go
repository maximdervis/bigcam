package routes

import (
	"server/src/controllers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	gymController controllers.GymController
}

func NewRoute(gymController controllers.GymController) Routes {
	return Routes{gymController}
}

func (r *Routes) Route(rg *gin.RouterGroup) {
	router := rg.Group("gyms")
	router.POST("/", r.gymController.CreateGym)
	router.GET("/:gym_id", r.gymController.GetGym)
}
