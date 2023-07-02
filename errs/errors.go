package errs

import (
	"fmt"
)

const (
	nilString  = "nil"
	nilMessage = "success"
)

type Error interface {
	Code() int32
	Error() string
}

type KonoError struct {
	ErrCode int32
	ErrMsg  string
	cause   error
}

func (e *KonoError) Error() string {
	if e == nil {
		return nilString
	}
	return fmt.Sprintf("%v:%s", e.ErrCode, e.ErrMsg)
}

func (e *KonoError) Code() int32 {
	if e == nil {
		return 0
	}
	return e.ErrCode
}

func (e *KonoError) ResetMsg(msg string) *KonoError {
	if e != nil {
		return New(e.Code(), msg)
	}
	return e
}

func (e *KonoError) WithCause(err error) *KonoError {
	if e != nil {
		e.cause = err
		newMsg := fmt.Sprintf("%v,caused by %v", e.ErrMsg, err.Error())
		return New(e.Code(), newMsg)
	}
	return e
}

func New(code int32, message string) *KonoError {
	return &KonoError{ErrCode: code, ErrMsg: message}
}

func IsEqual(a, b Error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Code() == b.Code()
}
