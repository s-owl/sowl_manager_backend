package controllers

import (
	"strings"
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/models"
)

// userSignupLogic 유저 회원 가입 로직
func userSignupLogic(c context.Context, userData *models.UserData) (*models.User, error) {
	if !strings.HasSuffix(userData.Email, "office.skhu.ac.kr") {
		return nil, EmailError
	}

	if len(userData.Password) < 8 {
		return nil, PasswordError
	}

	if userData.Password != userData.PasswordCheck {
		return nil, PasswordCheckError
	}

	if userData.Name == "" {
		return nil, NameError
	}

	userInfo := models.User{
		Email:    userData.Email,
		Name:     userData.Name,
		Nickname: userData.Nickname,
		Contact:  userData.Contact,
	}

	authClient := firebaseapp.App().Auth

	// 에러가 발생하지 않으면 이 이메일로 존재하는 계정이 있다는 증명이다.
	_, err := authClient.GetUserByEmail(c, userData.Email)
	if err == nil {
		return nil, EmailDuplicationError
	}
	err = nil

	params := (&auth.UserToCreate{}).
		Email(userData.Email).
		EmailVerified(false).
		Password(userData.Password).
		DisplayName(userData.Name).
		Disabled(false)
	_, err = authClient.CreateUser(c, params)
	if err != nil {
		return nil, UserCreateError
	}

	_, _, err = userInfo.CreateUser(c)
	if err != nil {
		return nil, UserCreateError
	}

	return &userInfo, nil
}
