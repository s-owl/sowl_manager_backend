package utils

import (
	"net/http"
)

// GinJsonMarshalError Gin이 Json을 처리하는 과정에서 문제가 발생하면 사용한다.
func GinJSONMarshalError(internal error) error {
	return &internalError {
		"E1",
		"Not Json Format",
		http.StatusBadRequest,
		internal,
	}
}
