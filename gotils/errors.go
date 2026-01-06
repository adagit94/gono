package gotils

import (
	"fmt"
)

const (
	MethodNotRegisteredCode = iota
	RouteNotRegisteredCode
)

type CodeError struct {
	Code    int
	Message string
}

func (err *CodeError) Error() string {
	return fmt.Sprintf("[%d]: %s", err.Code, err.Message)
}
