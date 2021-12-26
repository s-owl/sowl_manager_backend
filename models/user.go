package models

import (
	"context"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/utils"
)

const userCollection = "users"

// UserSignupInput - 사용자 회원가입 데이터
type UserSignupInput struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	PasswordCheck string `json:"passwordCheck" binding:"required"`
	Name          string `json:"name" biding:"required"`
	Nickname      string `json:"nickname"`
	Contact       string `json:"contact"`
}

// Validate - 사용자 회원 가입 데이터를 검증한다.
func (userInput *UserSignupInput) Validate() error {
	if !strings.HasSuffix(userInput.Email, "office.skhu.ac.kr") {
		return UserEmailError
	}

	if len(userInput.Password) < 8 {
		return UserPasswordError
	}

	if userInput.Password != userInput.PasswordCheck {
		return UserPasswordCheckError
	}

	if userInput.Name == "" {
		return UserNameError
	}

	return nil
}

// User - Firestore에 저장할 사용자 데이터와 Auth와 통신할 데이터
type User struct {
	userRecord *auth.UserRecord `json:"-"`
	Email      string           `json:"email"`
	Name       string           `json:"name"`
	Nickname   string           `json:"nickname"`
	Contact    string           `json:"contact"`
	Groups     []string         `json:"groups"`
}

// CreateUser - 사용자 데이터를 입력받아 Auth, Firestore에 생성
func RegisterUser(context context.Context, userInput *UserSignupInput) (user *User, err error) {
	firestoreClient := firebaseapp.App().Firestore
	authClient := firebaseapp.App().Auth
	user = new(User)

	// UserSignupInput에서 데이터 추출
	user.Email = userInput.Email
	user.Name = userInput.Name
	user.Nickname = userInput.Nickname
	user.Contact = userInput.Contact

	// Auth에 회원 추가 시도
	user2create := (&auth.UserToCreate{}).
		Email(userInput.Email).
		EmailVerified(false).
		Password(userInput.Password).
		DisplayName(userInput.Name).
		Disabled(false)
	user.userRecord, err = authClient.CreateUser(context, user2create)
	if err != nil {
		user = nil
		err = utils.FirebaseAuthError(err)
		return
	}

	// Firestore에 저장 시도
	_, err = firestoreClient.Collection(userCollection).Doc(user.Email).Create(context, user)
	if err != nil {
		// firebase auth에서 삭제 시도를 하고 실패 시 gin에다 로깅한다.
		if authErr := authClient.DeleteUser(context, user.userRecord.UID); authErr != nil {
			if ginc, ok := context.(*gin.Context); ok {
				ginc.Error(authErr)
			}
		}

		user = nil
		err = utils.FirestoreError(err)

		return
	}

	return
}
