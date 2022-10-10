package gerror

import (
	"errors"
	"gameutil/errors/gcode"
	"runtime"
	"strings"
)

const (
	// Filtering key for current error module paths.
	stackFilterKeyLocal = "gamesrv/gerror/gerror"
)

var (
	// goRootForFilter is used for stack filtering in development environment purpose.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.Replace(goRootForFilter, "\\", "/", -1)
	}
}

type Error struct {
	error error
	stack stack
	text  string
	code  gcode.Code
}

func (err *Error) Error() string {
	if err == nil {
		return ""
	}

	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}

	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}

		errStr += err.error.Error()
	}

	return errStr
}

func (err *Error) Cause() error {
	if err == nil {
		return nil
	}

	loop := err
	for loop != nil {
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				loop = e
			} else if e, ok := loop.error.(ICause); ok {
				return e.Cause()
			} else {
				return loop.error
			}
		} else {
			return errors.New(loop.text)
		}
	}

	return nil
}

func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}

	return err.error
}
