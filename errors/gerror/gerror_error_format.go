package gerror

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
)

// Format formats the frame according to the fmt.Formatter interface.
//
// %v, %s   : Print all the error string;
// %-v, %-s : Print current level error string;
// %+s      : Print full stack error list;
// %+v      : Print the error string and full stack error list;
func (err *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			if err.text != "" {
				_, _ = io.WriteString(s, err.text)
			} else {
				_, _ = io.WriteString(s, err.Error())
			}
		case s.Flag('+'):
			if verb == 's' {
				_, _ = io.WriteString(s, err.Stack())
			} else {
				_, _ = io.WriteString(s, err.Error()+"\n"+err.Stack())
			}
		default:
			_, _ = io.WriteString(s, err.Error())
		}
	}
}

func formatSubStack(st stack, buffer *bytes.Buffer) {
	if st == nil {
		return
	}

	index := 1
	space := " "
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if strings.Contains(file, stackFilterKeyLocal) {
				continue
			}

			if strings.Contains(file, "<") {
				continue
			}

			// Ignore GO ROOT paths.
			if goRootForFilter != "" &&
				len(file) >= len(goRootForFilter) &&
				file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}

			// Graceful indent.
			if index > 9 {
				space = " "
			}
			buffer.WriteString(fmt.Sprintf(
				"   %d).%s%s\n    \t%s:%d\n",
				index, space, fn.Name(), file, line,
			))
			index++
		}
	}

}
