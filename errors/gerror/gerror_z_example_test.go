package gerror_test

import (
	"errors"
	"fmt"
	"gameutil/errors/gcode"
	"gameutil/errors/gerror"
)

func ExampleNewCode() {
	err := gerror.NewCode(gcode.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.Code(err))

	// Output:
	// My Error
	// 10000
}

func ExampleNewCodef() {
	err := gerror.NewCodef(gcode.New(10000, "", nil), "It's %s", "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.Code(err).Code())

	// Output:
	// It's My Error
	// 10000
}

func ExampleWrapCode() {
	err1 := errors.New("permission denied")
	err2 := gerror.WrapCode(gcode.New(10000, "", nil), err1, "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.Code(err2).Code())

	// Output:
	// Custom Error: permission denied
	// 10000
}

func ExampleWrapCodef() {
	err1 := errors.New("permission denied")
	err2 := gerror.WrapCodef(gcode.New(10000, "", nil), err1, "It's %s", "Custom Error")
	fmt.Println(err2.Error())
	fmt.Println(gerror.Code(err2).Code())

	// Output:
	// It's Custom Error: permission denied
	// 10000
}

func ExampleNew() {
	test2()

	// Output:
	// Unordered output
}

func test2() {
	panic1()
}

func panic1() {
	defer func() {
		if err := recover(); err != nil {
			err := gerror.Newf("msg %s", err)
			//
			if stack := gerror.Stack(err); stack != "" {
				fmt.Println("\nStack: \n" + stack)
			}
		}
	}()

	panic("err 发生错误")
}
