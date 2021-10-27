package apperrors

import "fmt"

// UserError - User API 관련 Error
type UserError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

func (e *UserError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.Message)
}

var (
	UserCreateError = &UserError{
		ErrorCode: "U1",
		Message:   "Error creating new user",
	}

	EmailDuplicationError = &UserError{
		ErrorCode: "U2",
		Message:   "This email is already used",
	}

	EmailError = &UserError{
		ErrorCode: "U3",
		Message:   "Email domain must be office.skhu.ac.kr",
	}

	NameError = &UserError{
		ErrorCode: "U4",
		Message:   "Display name must not be empty",
	}

	PasswordCheckError = &UserError{
		ErrorCode: "U5",
		Message:   "Password and password check is not match",
	}

	PasswordError = &UserError{
		ErrorCode: "U6",
		Message:   "Invalid password format",
	}

	NicknameError = &UserError{
		ErrorCode: "U7",
		Message:   "Invalid nickname format",
	}

	ContactError = &UserError{
		ErrorCode: "U8",
		Message:   "This phone number is already used",
	}
)
