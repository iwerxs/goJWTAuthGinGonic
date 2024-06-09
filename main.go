package main

import (
	routes "goJWTAuthGinGonic/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")

	if port==""{
		port="8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// two api
	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted api1"})
	})
	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"access granted api2"})
	})
	router.Run(":" + port)
}