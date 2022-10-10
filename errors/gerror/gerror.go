package gerror

import "gameutil/errors/gcode"

// ICode is the interface for Code feature.
type ICode interface {
	Error() string
	Code() gcode.Code
}

// IStack is the interface for Stack feature.
type IStack interface {
	Error() string
	Stack() string
}

// ICause is the interface for Cause feature.
type ICause interface {
	Error() string
	Cause() error
}

// IUnwrap is the interface for Unwrap feature.
type IUnwrap interface {
	Error() string
	Unwrap() error
}
