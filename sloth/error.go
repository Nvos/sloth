package sloth

import (
	"bytes"
	"fmt"
)

type Code int

type Error struct {
	// For internal use only, should not be exposed out of service
	Op string
	Err error
	InternalMessage string

	// Rest
	HumanizedMessage string
	Code    Code
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		_, _ = fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		_, _ = buf.WriteString(e.Err.Error())
	}

	if e.InternalMessage != "" {
		_, _ = fmt.Fprintf(&buf, ", %s", e.InternalMessage)
	}

	return buf.String()
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Wrap(err error) *Error {
	return &Error{
		Err: err,
	}
}

func (e *Error) WithOp(op string) *Error {
	e.Op = op
	return e
}

func WrapInternal(err error, op string, message string) *Error {
	return &Error{
		Err: err,
		Op: op,
		InternalMessage: message,
	}
}

func (e *Error) WithInternalMessage(message string) *Error {
	e.InternalMessage = fmt.Sprintf(message)
	return e
}

func (e *Error) WithExternal(code Code, message string) *Error {
	e.Code = code
	e.HumanizedMessage = message
	return e
}
