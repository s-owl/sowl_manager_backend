package models

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
)

const userCollection = "users"

// UserData - 사용자 회원가입 데이터
type UserData struct {
	Email         string `json:"email" binding:"required,email,endswith=skhu.ac.kr"`
	Password      string `json:"password" binding:"required"`
	PasswordCheck string `json:"passwordCheck" binding:"required"`
	Name          string `json:"name" biding:"required"`
	Nickname      string `json:"nickname"`
	Contact       string `json:"contact"`
}

// User - Firestore에 저장할 사용자 데이터
type User struct {
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	Nickname string   `json:"nickname"`
	Contact  string   `json:"contact"`
	Groups   []string `json:"groups"`
}

// CreateUser - 사용자 데이터를 Firestore에 생성
func (user *User) CreateUser(context context.Context) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	client := firebaseapp.App().Firestore
	doc, wr, err := client.Collection(userCollection).Add(context, user)
	return doc, wr, err
}
