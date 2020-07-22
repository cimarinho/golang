package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//v1 := router.Group("/v1")
	//{
	//	//v1.POST("/create", service.http.create)
	//	//v1.POST("/read", submitEndpoint)
	//	//v1.POST("/update", readEndpoint)
	//}

	router.Run(":8080")
}

