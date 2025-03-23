package routes

import (
	"github.com/gin-gonic/gin"
	"server/src/controllers"
)

type UserRoute struct {
	userController controllers.UserController
}

func NewUserRoute(userController controllers.UserController) UserRoute {
	return UserRoute{userController}
}

func (r *UserRoute) Route(rg *gin.RouterGroup) {
	router := rg.Group("user")
	router.POST("/sign-up", r.userController.SignUp)
	router.POST("/sign-in", r.userController.SignIn)
	router.PUT("/update", r.userController.UpdateUser)
	router.GET("/get", r.userController.GetUser)
}
