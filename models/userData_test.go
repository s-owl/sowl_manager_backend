package models

import (
	"context"
	"fmt"
	"testing"

	"github.com/s-owl/sowl_manager_backend/firebaseapp"
)

// 테스트용 유저 데이터
var user = User{
	Email:    "test@office.skhu.ac.kr",
	Name:     "tester",
	Nickname: "testNick",
	Contact:  "010-0000-0000",
}

// TestCreateUser
func TestCreateUser(t *testing.T) {
	context := context.Background()
	firebaseapp.InitFirebaseApp(context)

	doc, wr, err := user.CreateUser(context)
	fmt.Println(doc, wr, err)
}
