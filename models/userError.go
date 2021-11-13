package models

import (
	"net/http"
	"github.com/s-owl/sowl_manager_backend/utils"
)

var (
	UserEmailError = &utils.AppError{
		ErrorCode: "U1",
		Message:   "Email domain must be office.skhu.ac.kr",
		StatusCode: http.StatusBadRequest,
	}

	UserNameError = &utils.AppError{
		ErrorCode: "U2",
		Message:   "Display name must not be empty",
		StatusCode: http.StatusBadRequest,
	}

	UserPasswordCheckError = &utils.AppError{
		ErrorCode: "U3",
		Message:   "Password and password check is not match",
		StatusCode: http.StatusBadRequest,
	}

	UserPasswordError = &utils.AppError{
		ErrorCode: "U4",
		Message:   "Invalid password format",
		StatusCode: http.StatusBadRequest,
	}

	UserNicknameError = &utils.AppError{
		ErrorCode: "U5",
		Message:   "Invalid nickname format",
		StatusCode: http.StatusBadRequest,
	}

	UserContactError = &utils.AppError{
		ErrorCode: "U6",
		Message:   "This phone number is already used",
		StatusCode: http.StatusBadRequest,
	}
)
