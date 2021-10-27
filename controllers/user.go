package controllers

import (
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/apperrors"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/models"
)

// UserController - /user 라우팅 설정
func UserController(router *gin.RouterGroup) {
	r := router.Group("/user")
	{
		r.POST("/signup", userSignup)
	}
}

func userSignup(c *gin.Context) {
	userData := models.UserData{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailDomain := strings.Split(userData.Email, "@")
	if !strings.HasSuffix(emailDomain[1], "office.skhu.ac.kr") {
		c.AbortWithStatusJSON(http.StatusBadRequest, apperrors.EmailError)
		return
	}

	if len(userData.Password) < 8 {
		c.AbortWithStatusJSON(http.StatusBadRequest, apperrors.PasswordError)
		return
	}

	if userData.Password != userData.PasswordCheck {
		c.AbortWithStatusJSON(http.StatusBadRequest, apperrors.PasswordCheckError)
		return
	}

	if userData.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, apperrors.NameError)
		return
	}

	userInfo := models.User{
		Email:    userData.Email,
		Name:     userData.Name,
		Nickname: userData.Nickname,
		Contact:  userData.Contact,
	}

	_, _, err := userInfo.CreateUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apperrors.UserCreateError)
	}

	authClient := firebaseapp.App().Auth
	params := (&auth.UserToCreate{}).
		Email(userData.Email).
		EmailVerified(false).
		Password(userData.Password).
		DisplayName(userData.Name).
		Disabled(false)
	user, err := authClient.CreateUser(c, params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apperrors.UserCreateError)
		return
	}

	log.Printf("Successfully created user: %v", user)
}
