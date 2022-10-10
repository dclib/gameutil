package gerror

import "runtime"

type stack []uintptr

const (
	MAX_STACK_DEPTH = 64
)

func Cause(err error) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(ICause); ok {
		return e.Cause()
	}

	if e, ok := err.(IUnwrap); ok {
		return Cause(e.Unwrap())
	}

	return err
}

func Stack(err error) string {
	if err == nil {
		return ""
	}

	if e, ok := err.(IStack); ok {
		return e.Stack()
	}

	return err.Error()
}

func callers(skip ...int) stack {
	var (
		pcs [MAX_STACK_DEPTH]uintptr
		n   = 3
	)

	if len(skip) > 0 {
		n += skip[0]
	}

	return pcs[:runtime.Callers(n, pcs[:])]
}
