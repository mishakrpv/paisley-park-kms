package util

import (
	"errors"

	"go.uber.org/zap"
)

type HttpErr struct {
	Status int
	Err    error
}

func (e *HttpErr) Error() string { return string(rune(e.Status)) + ": " + e.Err.Error() }

func (e *HttpErr) Unwrap() error { return e.Err }

func (e *HttpErr) Timeout() bool {
	t, ok := e.Err.(interface{ Timeout() bool })
	return ok && t.Timeout()
}

func HandleErr(status int, msg string) *HttpErr {
	zap.L().Error("An error occured", zap.Error(errors.New(msg)))
	return &HttpErr{Status: status, Err: errors.New(msg)}
}
