package controllers

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/s-owl/sowl_manager_backend/models"
)

// groupRegistLogic 그룹 등록 로직
func groupRegistLogic(c context.Context, groupInput *models.GroupRegistInput) (*firestore.WriteResult, error) {
	if err := groupInput.Validate(); err != nil {
		return nil, err
	}

	wr, err := models.RegistGroup(c, groupInput)
	if err != nil {
		return nil, err
	}

	return wr, err
}