package controllers

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/models"
	"github.com/s-owl/sowl_manager_backend/utils"
)

// UserController - /user 라우팅 설정
func UserController(router *gin.RouterGroup) {
	r := router.Group("/user")
	{
		r.POST("/signup", userSignup)
	}
}

func userSignup(c *gin.Context) {
	var err error = nil
	var user *models.User
	userData := models.UserData{}

	if err = c.ShouldBindJSON(&userData); err == nil {
		user, err = userSignupLogic(c, &userData)
	} else {
		err = utils.GinJSONMarshalError(err)
	}

	if err != nil {
		err = fmt.Errorf("UserSignup: %w", err)
		utils.AbortWithHTTPError(c, err)
		return
	}

	log.Printf("Successfully created user: %v", user)
}
