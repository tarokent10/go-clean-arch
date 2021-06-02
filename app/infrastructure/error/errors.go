package errors

import (
	"fmt"
	"picture-go-app/infrastructure/error/codes"

	"github.com/pkg/errors"
)

type appError struct {
	code codes.Code
	err  error
}

func (e appError) Error() string {
	return fmt.Sprintf("Codes: %s, Msg %s", e.code, e.err)
}

// Errorf returns an error containing an error code and a description;
// Errorf returns nil if c is OK.
func Errorf(c codes.Code, format string, a ...interface{}) error {
	if c == codes.OK {
		return nil
	}
	return appError{
		code: c,
		err:  errors.Errorf(format, a...),
	}
}

// Code returns the error code for err if it was produced by this system.
// Otherwise, it returns codes.SystemError.
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if e, ok := err.(appError); ok {
		return e.code
	}
	return codes.Unknown
}

func StackTrace(err error) string {
	if e, ok := err.(appError); ok {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
