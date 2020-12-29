package main

import (
	"github.com/LuanSapelli/golang-gcm-cipher/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	api := router.Group("/gcm-cipher/api")
	api.GET("/health-check", controllers.HealthCheck)

	//GCM encrypt && decrypt endpoints
	api.POST("/encrypt-data", controllers.EncryptData)
	api.POST("/decrypt-data", controllers.DecryptData)

	err := router.Run(":80")
	if err != nil {
		log.Printf("Router error - %v", err)
	}
}
