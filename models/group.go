package models

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/utils"
)

const groupCollection = "groups"

// GroupRegistInput - 그룹 등록 데이터
type GroupRegistInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Permission  bool `json:"permission"`
}

// Validate - 그룹 등록 데이터를 검증한다.
func (group *GroupRegistInput) Validate() error {
	if (group.Name == "") {
		return GroupNameError
	}

	if (group.Description == "") {
		return GroupDescriptionError
	}

	if (group.Category == "") {
		return GroupCategoryError
	}

	return nil
}

// RegistGroup - 그룹의 데이터를 입력받아 Firestore에 생성
func RegistGroup(context context.Context, groupInput *GroupRegistInput) (*firestore.WriteResult, error) {
	firestoreClient := firebaseapp.App().Firestore

	// Firestore에 저장 시도
	wr, err := firestoreClient.Collection(groupCollection).Doc(groupInput.Name).Create(context, groupInput)
	if err != nil {
		err = utils.FirestoreError(err)

		return nil, err
	}

	return wr, err
}