package thetaerror

import (
	"bytes"
	"fmt"
	"github.com/WolffunService/theta-shared-common/thetalog"
)

type Error struct {
	// Machine-readable error code.
	Code int

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		_, _ = fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
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

func New(code int, message string, op string, err error) error {
	return &Error{Code: code, Message: message, Op: op, Err: err}
}

func NewWithLog(code int, message string, op string, err error) error {
	e := New(code, message, op, err)
	thetalog.Err(e)
	return e
}

func ErrorCode(err error) int {
	if err == nil {
		return 0
	} else if e, ok := err.(*Error); ok && e.Code != 0 {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return ErrorInternal
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	} else if ok && e.Message == "" {
		if msg, defined := errorText[e.Code]; defined {
			return msg
		}
	}

	return "An internal error has occurred. Please contact technical support."
}
