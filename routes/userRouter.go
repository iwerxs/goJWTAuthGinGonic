package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/iwerxs/goJWTAuthGinGonic/controllers"
	"github.com/iwerxs/goJWTAuthGinGonic/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}