package controllers

import (
	"net/http"
	"github.com/s-owl/sowl_manager_backend/utils"
)

var (
	UserCreateError = &utils.AppError{
		ErrorCode: "U1",
		Message:   "Error creating new user",
		StatusCode: http.StatusInternalServerError,
	}

	EmailDuplicationError = &utils.AppError{
		ErrorCode: "U2",
		Message:   "This email is already used",
		StatusCode: http.StatusBadRequest,
	}

	EmailError = &utils.AppError{
		ErrorCode: "U3",
		Message:   "Email domain must be office.skhu.ac.kr",
		StatusCode: http.StatusBadRequest,
	}

	NameError = &utils.AppError{
		ErrorCode: "U4",
		Message:   "Display name must not be empty",
		StatusCode: http.StatusBadRequest,
	}

	PasswordCheckError = &utils.AppError{
		ErrorCode: "U5",
		Message:   "Password and password check is not match",
		StatusCode: http.StatusBadRequest,
	}

	PasswordError = &utils.AppError{
		ErrorCode: "U6",
		Message:   "Invalid password format",
		StatusCode: http.StatusBadRequest,
	}

	NicknameError = &utils.AppError{
		ErrorCode: "U7",
		Message:   "Invalid nickname format",
		StatusCode: http.StatusBadRequest,
	}

	ContactError = &utils.AppError{
		ErrorCode: "U8",
		Message:   "This phone number is already used",
		StatusCode: http.StatusBadRequest,
	}
)
