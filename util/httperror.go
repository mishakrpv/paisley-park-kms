package util

import (
	"errors"
)

type HttpError struct {
	status int
	msg    string
}

func NewHttpError(status int, msg string) *HttpError {
	return &HttpError{status: status, msg: msg}
}

func (e *HttpError) Status() int { return e.status }

func (e *HttpError) Error() string { return e.msg }

func (e *HttpError) HttpError() (int, string) { return e.status, e.msg }

func (e *HttpError) Unwrap() error { return errors.New(e.msg) }
