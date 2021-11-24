package controllers

import (
	"fmt"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/email"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
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
	authClient := firebaseapp.App().Auth

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

	actionCodeSettings := &auth.ActionCodeSettings{
		URL: "http://localhost:8080/api/user/signup",
		HandleCodeInApp: false,
	}

	verifyLink, err := authClient.EmailVerificationLinkWithSettings(c, user.Email, actionCodeSettings)
	if err != nil {
		err = fmt.Errorf("VerifyEmailLink: %w", err)
		utils.VerifyLinkError(err)
		return
	}

	fmt.Println(user.Email, verifyLink)
	err = email.SendMail(user.Email, verifyLink)
	if err != nil {
		err = fmt.Errorf("SendEmail: %w", err)
		utils.SendEmailError(err)
		return
	}

	log.Printf("Successfully created user: %v", user)
}
