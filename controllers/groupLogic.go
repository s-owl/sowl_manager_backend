package controllers

import (
	"context"

	"github.com/s-owl/sowl_manager_backend/models"
)

// groupRegistLogic 그룹 등록 로직
func groupRegisterLogic(c context.Context, groupInput *models.GroupRegisterInput, userEmail string) (*models.Group, error) {
	if err := groupInput.Validate(); err != nil {
		return nil, err
	}

	groupInfo, err := models.RegisterGroup(c, groupInput, userEmail)
	if err != nil {
		return nil, err
	}

	return groupInfo, nil
}
