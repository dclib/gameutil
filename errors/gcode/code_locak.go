package gcode

import "fmt"

type localCode struct {
	code    int
	message string
	detail  interface{}
}

func (c localCode) Code() int {
	return c.code
}

func (c localCode) Message() string {
	return c.message
}

func (c localCode) Detail() interface{} {
	return c.detail
}

func (c localCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}

	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}

	return fmt.Sprintf(`%d`, c.code)
}
