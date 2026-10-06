package main

import (
	"errors"
	"fmt"
)

type MyTypedError struct {
	Value int
}

func (e *MyTypedError) Error() string {
	return "MyTypedError error"
}

var SpecificErr = &MyTypedError{Value: 999}

func do() error {
	return SpecificErr
}

func main() {
	err := do()
	if err != nil {
		err = fmt.Errorf("wrap error: %w", err)

		typedErr, ok := err.(*MyTypedError)
		fmt.Println("typedErr:", typedErr, "ok:", ok)

		typedErrorExample := &MyTypedError{}
		fmt.Printf("errors.Is(&MyTypedError{}): %v\n", errors.Is(err, typedErrorExample))
		fmt.Printf("errors.Is(&MyTypedError{Value: 999}): %v\n", errors.Is(err, SpecificErr))
		fmt.Printf("errors.As(&MyTypedError{}): %v\n", errors.As(err, &typedErrorExample))
		fmt.Printf("got err: %+v\n", *typedErrorExample)
	}
}
