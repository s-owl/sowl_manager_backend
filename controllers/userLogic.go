package controllers

import (
	"context"

	"github.com/s-owl/sowl_manager_backend/models"
)

// userSignupLogic 유저 회원 가입 로직
func userSignupLogic(c context.Context, userInput *models.UserSignupInput) (*models.User, error) {
	if err := userInput.Validate(); err != nil {
		return nil, err
	}

	userInfo := models.User{}
	err := userInfo.CreateUser(c, userInput)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
