package utils

import "fmt"

/*
internalError 다른 원인으로 발생하는 에러를 정의한다.
함수에서만 생성 가능하게 비공개처리합니다.
*/
type internalError struct {
	ErrorCode  string `json:"ErrorCode"`
	Message    string `json:"Message"`
	StatusCode int    `json:"-"`
	internal   error
}

// GetStatusCode HttpError Interface에서 필요한 상태 코드를 출력한다.
func (e *internalError) GetStatusCode() int {
	return e.StatusCode
}

// Error error interface로 출력할 문자열
func (e *internalError) Error() string {
	return fmt.Sprintf("%s: %s, %s", e.ErrorCode, e.Message, e.internal)
}

// Unwrap Unwrap 기능을 활용해 내부 에러를 확인할 수 있게 한다.
func (e *internalError) Unwrap() error {
	return e.internal
}
