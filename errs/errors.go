package errs

import (
	"fmt"
)

const (
	nilString  = "nil"
	nilMessage = "success"
)

type KonoError struct {
	ErrCode int32
	ErrMsg  string
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

func (e *KonoError) ResetErrMsg(msg string) *KonoError {
	if e != nil {
		return New(e.Code(), msg)
	}
	return e
}

func (e *KonoError) JoinErrMsg(err Error) *KonoError {
	if e != nil {
		newMsg := fmt.Sprintf("%v|%v", e.ErrMsg, err.Error())
		return New(e.Code(), newMsg)
	}
	return e
}

func New(code int32, message string) *KonoError {
	err := &KonoError{
		ErrCode: code,
		ErrMsg:  message,
	}
	return err
}

type Error interface {
	Code() int32
	Error() string
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
