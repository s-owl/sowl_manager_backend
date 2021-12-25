package models

import (
	"context"
	"time"

	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/utils"
)

const groupCollection = "groups"

// GroupRegistInput - 그룹 등록 데이터
type GroupRegistInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

// Group - Firestore에 저장할 그룹 데이터
type Group struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Admin       string    `json:"admin"`
	Member      []string  `json:"member"`
	Permission  bool      `json:"permission"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Validate - 그룹 등록 데이터를 검증한다.
func (group *GroupRegistInput) Validate() error {
	if group.Name == "" || group.Description == "" || group.Category == "" {
		return GroupFieldNullError
	}

	return nil
}

// RegistGroup - 그룹의 데이터를 입력받아 Firestore에 생성
func RegistGroup(context context.Context, groupInput *GroupRegistInput, userEmail string) (group *Group, err error) {
	firestoreClient := firebaseapp.App().Firestore
	group = new(Group)

	// GroupRegistInput에서 데이터 추출
	group.Name = groupInput.Name
	group.Description = groupInput.Description
	group.Category = groupInput.Category
	group.Admin = userEmail
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	// Firestore에 저장 시도
	_, err = firestoreClient.Collection(groupCollection).Doc(group.Name).Create(context, group)
	if err != nil {
		group = nil
		err = utils.FirestoreError(err)

		return
	}

	return
}
