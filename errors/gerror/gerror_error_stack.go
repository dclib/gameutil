package gerror

import (
	"bytes"
	"fmt"
)

func (err *Error) Stack() string {
	if err == nil {
		return ""
	}

	var (
		loop   = err
		index  = 1
		buffer = bytes.NewBuffer(nil)
	)

	for loop != nil {
		buffer.WriteString(fmt.Sprintf("%d. %-v\n", index, loop))
		index++
		formatSubStack(loop.stack, buffer)
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				loop = e
			} else {
				buffer.WriteString(fmt.Sprintf("%d. %s\n", index, loop.error.Error()))
				index++
				break
			}
		} else {
			break
		}
	}

	return buffer.String()
}
