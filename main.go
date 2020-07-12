//Errors..
// Using unwrap

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	fmt.Println(err)

	//bar error
	baseErr := errors.Unwrap(err)
	fmt.Println(baseErr)

	//moo error
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	//cat error
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

	//no more errors
	baseErr = errors.Unwrap(baseErr)
	fmt.Println(baseErr)

}

func foo() error {
	return fmt.Errorf("foo is an error: %w", bar())
}

func bar() error {
	return fmt.Errorf("bar is an error: %w", moo())
}

func moo() error {
	return fmt.Errorf("moo is an error: %w", cat())
}

//base error
func cat() error {
	return errors.New("cat is an error")
}
