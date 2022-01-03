package utils

import "net/http"

func SendEmailError(intenal error) error {
	return &internalError{
		"E1",
		intenal.Error(),
		http.StatusInternalServerError,
		intenal,
	}
}

func VerifyLinkError(internal error) error {
	return &internalError{
		"E2",
		internal.Error(),
		http.StatusInternalServerError,
		internal,
	}
}
