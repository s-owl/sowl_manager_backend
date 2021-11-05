package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/controllers"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
)

func main() {

	router := gin.Default()

	// router 그룹 생성
	api := router.Group("/api")

	// router Test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	controllers.UserController(api)

	// firebase SDK Initialize
	firebaseapp.InitFirebaseApp(context.Background())

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":8080")
}
