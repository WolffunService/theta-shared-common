package thetanerr

import (
	"bytes"
	"fmt"
)

const defaultErrorMessage = "An internal error has occurred. Please contact technical support."

type ThetanError struct {
	// Machine-readable error code.
	Code ErrorCode

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error
}

// Error returns the string representation of the error message.
func (e *ThetanError) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		_, _ = fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise, print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != 0 {
			_, _ = fmt.Fprintf(&buf, "<%v> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func New(code ErrorCode) error {
	message := ErrMessageByCode(code)
	return &ThetanError{
		Code:    code,
		Message: message,
	}
}

func Wrap(err error, code ErrorCode) error {
	message := fmt.Sprintf("%s: %s", ErrMessageByCode(code), err.Error())
	return &ThetanError{
		Code:    code,
		Message: message,
	}
}

func NewCustomMessage(code ErrorCode, message string) error {
	return &ThetanError{
		Code:    code,
		Message: message,
	}
}

func ErrCode(err error) ErrorCode {
	if err == nil {
		return Success
	}

	if e, ok := err.(ErrorCode); ok {
		return e
	}

	if e, ok := err.(*ThetanError); ok && e.Code != 0 {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrCode(e.Err)
	}

	return ErrInternal
}

func ErrMessage(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(ErrorCode); ok {
		return e.Error()
	}

	if e, ok := err.(*ThetanError); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrMessage(e.Err)
	} else if ok && e.Message == "" {
		return ErrMessageByCode(e.Code)
	}

	return defaultErrorMessage
}

func ErrMessageByCode(code ErrorCode) string {
	if msg, defined := errMessage[code]; defined {
		return msg
	}

	return defaultErrorMessage
}
