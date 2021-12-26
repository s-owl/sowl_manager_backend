package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/email"
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

// userSignup godoc
// @Summary 유저 회원가입
// @Description 유저 회원가입 API
// @ID user-signup
// @Accept json
// @Produce json
// @Param newAdminForm body models.UserSignupInput true "유저 회원가입을 위한 양식"
// @Success 200 {object} models.InfoDTO
// @Failure 400 {object} models.ErrorDTO
// @Router /user/signup [post]
func userSignup(c *gin.Context) {
	var err error = nil
	var user *models.User
	userInput := models.UserSignupInput{}

	if err = c.ShouldBindJSON(&userInput); err == nil {
		user, err = userSignupLogic(c, &userInput)
	} else {
		err = utils.GinJSONMarshalError(err)
	}

	if err != nil {
		err = fmt.Errorf("UserSignup: %w", err)
		utils.AbortWithHTTPError(c, err)
		return
	}

	verifyLink, err := email.ExtractVerifyLink(c, userInput.Email)
	if err != nil {
		log.Printf("VerifyLink Error: %v", err)
		return
	}
	err = email.SendEmail(userInput.Email, verifyLink)
	if err != nil {
		log.Printf("SendEmail Error: %v", err)
		return
	}

	log.Printf("Successfully created user: %v", user)
}
