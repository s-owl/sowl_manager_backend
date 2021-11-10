package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/controllers"
	"github.com/s-owl/sowl_manager_backend/docs"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {

	// swagger 초기 설정
	docs.SwaggerInfo.Title = "Sowl-Manager RESTful API"
	docs.SwaggerInfo.Description = "RESTful API spec for Sowl-Manager Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	// swagger
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":8080")
}
