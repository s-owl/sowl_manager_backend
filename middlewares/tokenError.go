package middlewares

import (
	"net/http"

	"github.com/s-owl/sowl_manager_backend/utils"
)

var (
	UserTokenInvalidError = &utils.AppError{
		ErrorCode:  "T1",
		Message:    "User ID Token is not valid",
		StatusCode: http.StatusForbidden,
	}

	NoPermissionError = &utils.AppError{
		ErrorCode:  "T2",
		Message:    "User has no permission to access this resource or action",
		StatusCode: http.StatusForbidden,
	}
)
