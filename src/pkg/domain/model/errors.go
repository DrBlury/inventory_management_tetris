package domain

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

type ErrorType string

// Defines values for ErrorType.
const (
	BadRequest          ErrorType = "Bad Request"
	InternalServerError ErrorType = "Internal Server Error"
	NotFound            ErrorType = "Not Found"
	NotImplemented      ErrorType = "Not Implemented"
)

type Error struct {
	ErrorType     ErrorType
	Code          int
	FriendlyError string
	ErrorId       string
	Timestamp     time.Time
	DebugError    error
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling structured logging for Error.
func (e Error) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ErrorId", e.ErrorId)
	enc.AddInt("Code", e.Code)
	enc.AddString("Error", e.FriendlyError)
	enc.AddString("ErrorType", string(e.ErrorType))
	enc.AddTime("Timestamp", e.Timestamp)
	enc.AddString("DebugError", fmt.Sprintf("%+v", e.DebugError))
	return nil
}
