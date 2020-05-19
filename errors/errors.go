package errors

import (
	"fmt"
	"runtime"
	"strconv"
)

// AppError provides interface for application errors
type AppError interface {
	Error() string
	Caller() string
	ErrorStack() string
	StatusCode() int
}

type appError struct {
	err        error
	caller     string
	message    string
	statusCode int
}

// NewAppError create an app error
func NewAppError(message string, statusCode int, err error) AppError {
	caller := "UNKNOWN CALLER"
	pc, _, lineNumber, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller = details.Name() + ":" + strconv.Itoa(lineNumber)
	}

	return &appError{
		err:        err,
		caller:     caller,
		message:    message,
		statusCode: statusCode,
	}
}

func (err *appError) Caller() string {
	return err.caller
}

func (err *appError) Error() string {
	return err.message
}

func (err *appError) StatusCode() int {
	return err.statusCode
}

func (err *appError) ErrorStack() string {
	errorStack := "ErrorStack :-\n"

	var tempError error
	tempError = err
	for tempError != nil {
		errMessage := ""
		if newError, ok := tempError.(*appError); ok {
			errMessage = fmt.Sprintf("\t%s - %s", newError.Caller(), newError.Error())
			tempError = newError.err
		} else {
			errMessage = fmt.Sprintf("\tError: %s", tempError.Error())
			tempError = nil
		}
		errorStack += errMessage
	}
	return errorStack
}

// IsAppError checks if the error is appError and return the object
func IsAppError(err interface{}) (AppError, bool) {
	switch newError := err.(type) {
	case *appError:
		return newError, true
	default:
		return nil, false
	}
}
