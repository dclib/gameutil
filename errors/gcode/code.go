package gcode

type Code interface {
	Code() int
	Message() string
	Detail() interface{}
}

// ================================================================================================================
// Common error code definition.
// There are reserved internal error code by framework: code < 1000.
// ================================================================================================================

var (
	CodeNil = localCode{-1, "", nil}  // No error code specified.
	CodeOK  = localCode{0, "OK", nil} // It is OK.
)

func New(code int, message string, detail interface{}) Code {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

func WithCode(code Code, detail interface{}) Code {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}
