package models

import (
	"net/http"

	"github.com/s-owl/sowl_manager_backend/utils"
)

var (
	GroupFieldNullError = &utils.AppError{
		ErrorCode:  "G1",
		Message:    "Please input all the information",
		StatusCode: http.StatusBadRequest,
	}

	GroupKoreanNameError = &utils.AppError{
		ErrorCode:  "G2",
		Message:    "Group Korean Name can use only korean",
		StatusCode: http.StatusBadRequest,
	}

	GroupEnglishNameError = &utils.AppError{
		ErrorCode:  "G3",
		Message:    "Group English name can use only lower case",
		StatusCode: http.StatusBadRequest,
	}
)
