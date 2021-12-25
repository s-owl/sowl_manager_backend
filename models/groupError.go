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
)
