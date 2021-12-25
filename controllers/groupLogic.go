package controllers

import (
	"context"

	"github.com/s-owl/sowl_manager_backend/models"
)

// groupRegistLogic 그룹 등록 로직
func groupRegistLogic(c context.Context, groupInput *models.GroupRegistInput, userEmail string) (*models.Group, error) {
	if err := groupInput.Validate(); err != nil {
		return nil, err
	}

	groupInfo, err := models.RegistGroup(c, groupInput, userEmail)
	if err != nil {
		return nil, err
	}

	return groupInfo, nil
}
