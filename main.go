// Hands-on exercise #6 - Ninja Level 4
// Recreate the foo bar moo cat example. To do this, follow these general steps:
// create these functions, each of which returns a value of type error
// main calls foo
// foo
// foo calls bar
// bar
// bar calls moo
// moo
// moo calls cat
// cat
// cat calls nothing
// include the error information from any previous calls
// use errors.Unwrap() to access all of the individual errors
// name the variables that hold the unwrapped error
// fooErr
// barErr
// mooErr
// catErr
// baseErr
// this err will be nil
// tag v1.32.0
// video: 95 Hands On 6

package main

import (
	"errors"
	"fmt"
)

func main() {

	//New...
	//
	fmt.Println("Todds...")
	fooErr := foo()
	barErr := errors.Unwrap(fooErr)
	mooErr := errors.Unwrap(barErr)
	catErr := errors.Unwrap(mooErr)
	baseErr := errors.Unwrap(catErr)
	fmt.Printf("fooErr\t%s\n", fooErr)
	fmt.Printf("barErr\t%s\n", barErr)
	fmt.Printf("mooErr\t%s\n", mooErr)
	fmt.Printf("catErr\t%s\n", catErr)
	fmt.Printf("baseErr\t%v\n", baseErr)

	//Mine...
	//
	fmt.Println("Mine...")

	//foo error
	fooErr = foo()
	fmt.Println("foo error: ", fooErr)

	//bar error
	barErr = errors.Unwrap(fooErr)
	fmt.Println("bar error: ", barErr)

	//moo error
	mooErr = errors.Unwrap(barErr)
	fmt.Println("moo error: ", mooErr)

	//cat error
	catErr = errors.Unwrap(mooErr)
	fmt.Println("cat error: ", catErr)

	//no more errors
	baseErr = errors.Unwrap(catErr)
	fmt.Println("base error: ", baseErr)

	//Orig...
	//
	fmt.Println("Orig...")
	//foo error
	err := foo()
	fmt.Println(err)

	//bar error
	baseErr = errors.Unwrap(err)
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
	//call bar() before returning
	return fmt.Errorf("foo is an error: %w", bar())
}

func bar() error {
	//call moo() before returning
	return fmt.Errorf("bar is an error: %w", moo())
}

func moo() error {
	//call cat before returning
	return fmt.Errorf("moo is an error: %w", cat())
}

func cat() error {
	//create the base error
	return errors.New("cat is an error")
}
