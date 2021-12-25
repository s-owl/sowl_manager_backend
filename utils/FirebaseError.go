package utils

import (
	"net/http"
)

func FirebaseAuthError(internal error) error {
	return &internalError{
		"F1",
		internal.Error(),
		http.StatusInternalServerError,
		internal,
	}
}

func FirestoreError(internal error) error {
	return &internalError{
		"F2",
		"Database Error",
		http.StatusInternalServerError,
		internal,
	}
}
