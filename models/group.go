package models

import (
	"context"
	"fmt"
	"time"

	"github.com/s-owl/sowl_manager_backend/firebaseapp"
	"github.com/s-owl/sowl_manager_backend/utils"
)

const groupCollection = "groups"

// GroupRegistInput - 그룹 등록 데이터
type GroupRegistInput struct {
	KoreanName  string `json:"koreanName" binding:"required"`
	EnglishName string `json:"englishName" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

// Group - Firestore에 저장할 그룹 데이터
type Group struct {
	KoreanName  string    `json:"koreanName"`
	EnglishName string    `json:"englishName"`
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

	// 그룹 국문명이 한글인지 확인
	for _, value := range group.KoreanName {
		if value < 44032 || value > 55203 {
			fmt.Println(value)
			return GroupKoreanNameError
		}
	}

	// 그룹 영문명이 알파벳 소문자인지 확인
	for _, value := range group.EnglishName {
		if value < 97 || value > 122 {
			fmt.Println(value)
			return GroupEnglishNameError
		}
	}

	// 비어있는 필드가 있으면 오류 반환
	if group.KoreanName == "" || group.EnglishName == "" || group.Description == "" || group.Category == "" {
		return GroupFieldNullError
	}

	return nil
}

// RegistGroup - 그룹의 데이터를 입력받아 Firestore에 생성
func RegistGroup(context context.Context, groupInput *GroupRegistInput, userEmail string) (group *Group, err error) {
	firestoreClient := firebaseapp.App().Firestore
	group = new(Group)

	// GroupRegistInput에서 데이터 추출
	group.KoreanName = groupInput.KoreanName
	group.EnglishName = groupInput.EnglishName
	group.Description = groupInput.Description
	group.Category = groupInput.Category
	group.Admin = userEmail
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	// Firestore에 저장 시도
	_, err = firestoreClient.Collection(groupCollection).Doc(group.EnglishName).Create(context, group)
	if err != nil {
		group = nil
		err = utils.FirestoreError(err)

		return
	}

	return
}

// GetGroupByEnglishName - 그룹 영문명으로 해당 그룹 조회
func GetGroupByEnglishName(context context.Context, groupEnglishName string) (*Group, error) {
	var group Group
	firestoreClient := firebaseapp.App().Firestore

	groupData, err := firestoreClient.Collection(groupCollection).Doc(groupEnglishName).Get(context)
	if err != nil {
	}

	groupData.DataTo(&group)

	return &group, nil
}
