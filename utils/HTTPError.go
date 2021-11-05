package utils

import (
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
)

// HTTPError Gin 으로 내보낼 수 있는 에러를 정의합니다.
type HTTPError interface {
	GetStatusCode() int
	Error() string
}

/*
AbortWithHTTPError
HTTPError 인 에러면 정확한 상태코드와 메세지를 내보냅니다.
하지만 없다면 상태코드 500의 "Unexpected Error"로 출력해서 보냅니다.
*/
func AbortWithHTTPError(c *gin.Context, err error) {
	var httpError HTTPError
	if !errors.As(err, &httpError) {
		httpError = &internalError {
			"E0",
			"Unexptected Error",
			http.StatusInternalServerError,
			err,
		}
		err = httpError
	}

	c.AbortWithStatusJSON(httpError.GetStatusCode(), httpError)
	c.Error(err)
}
