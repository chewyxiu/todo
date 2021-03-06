package logic

import "errors"

var ApiError = func(statusCode int, errText string) APIError {
	return APIError{
		statusCode: statusCode,
		err:        errors.New(errText),
	}
}

type APIError struct {
	statusCode int
	err        error
}

func (a APIError) StatusCode() int {
	return a.statusCode
}

func (a APIError) Error() string {
	return a.err.Error()
}

