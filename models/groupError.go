package models

import (
	"net/http"

	"github.com/s-owl/sowl_manager_backend/utils"
)

var (
	GroupNameError = &utils.AppError{
		ErrorCode: "G1",
		Message:   "Display name must not be empty",
		StatusCode: http.StatusBadRequest,
	}

	GroupDescriptionError = &utils.AppError{
		ErrorCode: "G2",
		Message: "Display description must not be empty",
		StatusCode: http.StatusBadRequest,
	}

	GroupCategoryError = &utils.AppError{
		ErrorCode: "G3",
		Message: "Choose Category must not be empty",
		StatusCode: http.StatusBadRequest,
	}
)