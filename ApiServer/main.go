package main

import (
	"github.com/gin-gonic/gin"
	"goingo/ApiServer/routes"
)

func main() {
	router := gin.Default()
	router.GET("/sizeoptions", routes.GetSizeOptions)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
