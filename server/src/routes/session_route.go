package routes

import (
	"server/src/controllers"
	"server/src/middlewares"

	"github.com/gin-gonic/gin"
)

type SessionRoute struct {
	SessionController controllers.SessionController
}

func NewSessionRoute(sessionController controllers.SessionController) SessionRoute {
	return SessionRoute{sessionController}
}

func (r *SessionRoute) Route(rg *gin.RouterGroup) {
	router := rg.Group("sessions")
	router.Use(middlewares.IsAuthorized())
	router.GET("/list", r.SessionController.ListSessions)
	router.POST("/finish", r.SessionController.FinishSession)
	router.POST("/create", r.SessionController.CreateSession)
}
