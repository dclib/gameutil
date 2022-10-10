package gerror

import (
	"fmt"

	"github.com/dclib/gameutil/errors/gcode"

	"strings"
)

func NewCode(code gcode.Code, text ...string) error {
	return &Error{
		stack: callers(),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

func NewCodef(code gcode.Code, format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

func WrapCode(code gcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  strings.Join(text, ", "),
		code:  code,
	}
}

func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &Error{
		error: err,
		stack: callers(),
		text:  fmt.Sprintf(format, args...),
		code:  code,
	}
}

func Code(err error) gcode.Code {
	if err == nil {
		return gcode.CodeNil
	}

	if e, ok := err.(ICode); ok {
		return e.Code()
	}

	if e, ok := err.(IUnwrap); ok {
		return Code(e.Unwrap())
	}

	return gcode.CodeNil
}
