package main

import (
	"log"
	"os"

	routes "github.com/iwerxs/goJWTAuthGinGonic/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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