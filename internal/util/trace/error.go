package trace

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

const PrefixBadRequest = "BAD REQUEST"
const PrefixNotFound = "NOT FOUND"
const PrefixBadGateway = "BAD GATEWAY"
const PrefixTimeOut = "TIMEOUT"
const PrefixInternalServer = "INTERNAL SERVER ERROR"
const PrefixUnAuthorization = "UN AUTHORIZATION"
const PrefixForbidden = "FORBIDDEN"

type ErrTrace struct {
	Stack        error
	ApiCallError error
	Msg          string
}

func (err *ErrTrace) Error() string {
	return err.Stack.Error()
}

func StackTrace(err error) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          errTrace.Msg,
			ApiCallError: errTrace.ApiCallError,
		}
	} else {
		return &ErrTrace{
			Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:   PrefixInternalServer,
		}
	}
}

func ErrInternalServer(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixInternalServer + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixInternalServer
	if len(messages) > 0 {
		msg = PrefixInternalServer + " - " + strings.Join(messages, ". ")
	}
	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrBadRequest(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixBadRequest + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixBadRequest
	if len(messages) > 0 {
		msg = PrefixBadRequest + " - " + strings.Join(messages, ". ")
	} else {
		msg = PrefixBadRequest + " - " + err.Error()
	}

	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrNotFound(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixNotFound + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixNotFound
	if len(messages) > 0 {
		msg = PrefixNotFound + " - " + strings.Join(messages, ". ")
	}
	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrBadGateway(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixBadGateway + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixBadGateway
	if len(messages) > 0 {
		msg = PrefixBadGateway + " - " + strings.Join(messages, ". ")
	}
	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrTimeOut(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixTimeOut + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixTimeOut
	if len(messages) > 0 {
		msg = PrefixTimeOut + " - " + strings.Join(messages, ". ")
	}
	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrUnAuthorization(err error, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = PrefixUnAuthorization + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := PrefixUnAuthorization
	if len(messages) > 0 {
		msg = PrefixUnAuthorization + " - " + strings.Join(messages, ". ")
	}

	return &ErrTrace{
		Stack: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:   msg,
	}
}

func ErrorAsNotFound(err error) bool {
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		if strings.Contains(errTrace.Msg, PrefixNotFound) {
			return true
		}
	}

	return false
}

func ErrExternalApiCall(err error, prefix string, messages ...string) *ErrTrace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	var errTrace *ErrTrace
	if errors.As(err, &errTrace) {
		msg := errTrace.Msg
		if len(messages) > 0 {
			msg = prefix + " - " + strings.Join(messages, ". ")
		}

		return &ErrTrace{
			Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
			Msg:          msg,
			ApiCallError: errTrace.ApiCallError,
		}
	}

	msg := prefix
	if len(messages) > 0 {
		msg = prefix + " - " + strings.Join(messages, ". ")
	}

	return &ErrTrace{
		Stack:        fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		Msg:          msg,
		ApiCallError: err,
	}
}
