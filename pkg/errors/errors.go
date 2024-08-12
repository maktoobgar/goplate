package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"service/pkg/colors"
	"strings"
)

type (
	serverError struct {
		code    int64
		message string
		errMsg  string
		errors  any

		stack string
	}
)

const (
	_ int64 = iota
	// BadRequest 400
	InvalidStatus
	// NotFound 404
	NotFoundStatus
	// Unauthorized 401
	UnauthorizedStatus
	// InternalServerError 500
	UnexpectedStatus
	// MethodNotAllowed 405
	MethodNotAllowedStatus
	// Forbidden 403
	ForbiddenStatus
	// Timeout 408
	TimeoutStatus
	// Service Unavailable 503
	ServiceUnavailable
	// TooManyRequests 429
	TooManyRequests
)

var (
	httpErrors = map[int64]int{
		InvalidStatus:          http.StatusBadRequest,
		NotFoundStatus:         http.StatusNotFound,
		UnauthorizedStatus:     http.StatusUnauthorized,
		UnexpectedStatus:       http.StatusInternalServerError,
		MethodNotAllowedStatus: http.StatusMethodNotAllowed,
		ForbiddenStatus:        http.StatusForbidden,
		TimeoutStatus:          http.StatusRequestTimeout,
		ServiceUnavailable:     http.StatusServiceUnavailable,
		TooManyRequests:        http.StatusTooManyRequests,
	}
)

func (e serverError) Error() string {
	errCode := httpErrors[e.code]
	if errCode >= 500 {
		return fmt.Sprintf("%sCode: %d - Message: %s - Error: %s%s", colors.Red, errCode, e.message, e.errMsg, colors.Reset)
	}
	message := fmt.Sprintf("%sCode: %d - Message: %s - Error: %s%s", colors.Orange, errCode, e.message, e.errMsg, colors.Reset)
	if e.errMsg == "" {
		message = fmt.Sprintf("%sCode: %d - Message: %s%s", colors.Green, errCode, e.message, colors.Reset)
	}
	return message
}

func (e serverError) HasStackError() bool {
	return e.stack != ""
}

func (e serverError) GetStack() string {
	return e.stack
}

func (e serverError) GetErrorMessage() string {
	return e.errMsg
}

func CastError(err error) *serverError {
	if er, ok := err.(serverError); ok {
		return &er
	}

	return nil
}

// Returns httpErrorCode, message of it
func HttpError(err error) (code int, message string, errMsg string, errors any) {
	code = http.StatusInternalServerError
	errMsg = err.Error()
	message = errMsg
	errors = nil

	if er, ok := err.(serverError); ok {
		code = httpErrors[er.code]
		errMsg = er.errMsg
		message = er.message
		errors = er.errors
	}

	return
}

func IsServerError(err error) bool {
	if _, ok := err.(serverError); ok {
		return true
	}
	return false
}

// Creates a new error
func New(code int64, message string, errMsg string, errors ...any) error {
	var errs any = nil
	if len(errors) != 0 {
		errs = errors[0]
	}
	return serverError{
		code:    code,
		errMsg:  errMsg,
		message: message,
		errors:  errs,
		stack:   string(debug.Stack()),
	}
}

func IsContextDeadlineExceeded(err error) bool {
	return strings.Contains(err.Error(), "context deadline exceeded") || strings.Contains(err.Error(), "context canceled")
}
