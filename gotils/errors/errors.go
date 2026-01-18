package errors

import (
	"fmt"
)

const (
	MethodNotRegisteredCode = iota
	RouteNotRegisteredCode
	HandlerNotFoundCode
)

type CodeError struct {
	Code    int
	Message string
}

func (err *CodeError) Error() string {
	return fmt.Sprintf("[%d]: %s", err.Code, err.Message)
}

func Safely(f func(), onPanic func(v any)) {
	defer func() {
		if r := recover(); r != nil {
			onPanic(r)
		}
	}()

	f()
}
