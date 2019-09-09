package errors

import (
	"fmt"
	"runtime"
	"strconv"
)

type statusCode int

// constants to be used in Error as StatusCodes
const (
	StatusBadRequest                   statusCode = 400
	StatusUnauthorized                 statusCode = 401
	StatusPaymentRequired              statusCode = 402
	StatusForbidden                    statusCode = 403
	StatusNotFound                     statusCode = 404
	StatusMethodNotAllowed             statusCode = 405
	StatusNotAcceptable                statusCode = 406
	StatusProxyAuthRequired            statusCode = 407
	StatusRequestTimeout               statusCode = 408
	StatusConflict                     statusCode = 409
	StatusGone                         statusCode = 410
	StatusLengthRequired               statusCode = 411
	StatusPreconditionFailed           statusCode = 412
	StatusRequestEntityTooLarge        statusCode = 413
	StatusRequestURITooLong            statusCode = 414
	StatusUnsupportedMediaType         statusCode = 415
	StatusRequestedRangeNotSatisfiable statusCode = 416
	StatusExpectationFailed            statusCode = 417
	StatusTeapot                       statusCode = 418
	StatusMisdirectedRequest           statusCode = 421
	StatusUnprocessableEntity          statusCode = 422
	StatusLocked                       statusCode = 423
	StatusFailedDependency             statusCode = 424
	StatusTooEarly                     statusCode = 425
	StatusUpgradeRequired              statusCode = 426
	StatusPreconditionRequired         statusCode = 428
	StatusTooManyRequests              statusCode = 429
	StatusRequestHeaderFieldsTooLarge  statusCode = 431
	StatusUnavailableForLegalReasons   statusCode = 451

	StatusInternalServerError           statusCode = 500
	StatusNotImplemented                statusCode = 501
	StatusBadGateway                    statusCode = 502
	StatusServiceUnavailable            statusCode = 503
	StatusGatewayTimeout                statusCode = 504
	StatusHTTPVersionNotSupported       statusCode = 505
	StatusVariantAlsoNegotiates         statusCode = 506
	StatusInsufficientStorage           statusCode = 507
	StatusLoopDetected                  statusCode = 508
	StatusNotExtended                   statusCode = 510
	StatusNetworkAuthenticationRequired statusCode = 511
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
	statusCode statusCode
}

// NewAppError create an app error
func NewAppError(message string, statusCode statusCode, err error) AppError {
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
	return int(err.statusCode)
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
			errMessage = fmt.Sprintf("\tError: %s", err.Error())
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
