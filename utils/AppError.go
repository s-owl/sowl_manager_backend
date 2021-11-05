package utils

import (
	"fmt"
)

// AppError 어플리케이션(컨트롤러)에서 발생하는 에러
type AppError struct {
	ErrorCode  string `json:"ErrorCode"`
	Message    string `json:"Message"`
	StatusCode int    `json:"-"`
}

// Error Error interface를 통해 에러를 출력
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.Message)
}

// GetStatusCode HttpError interface의 동작에 필요한 http 상태 코드를 제공한다.
func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}
