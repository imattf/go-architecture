//Hands-on Exercise #1 - Ninja Level 3
// Create a background context.
// Create a child context with that background as a parent
// store two values in the child context
// key of type string
// key of type int
// pull the values out
// tag v1.13.0
// video: 67 Hands On 1

package main

import (
	"context"
	"fmt"
)

// type key int

// var userKey key = 0
// var admin key = 1

func main() {

	//create parent background context
	ctx := context.Background()

	//create child context w/ 2 values
	ctx = context.WithValue(ctx, "userID", 1234)

	//append another value to the child context
	ctx = context.WithValue(ctx, 1, "admin")

	//pull the values out
	val := ctx.Value("userID")
	if val == nil {
		fmt.Println("No value for key userID")
	} else {
		fmt.Println("userID =", val)
	}

	val = ctx.Value(1)
	if val == nil {
		fmt.Println("userID is not an admin")
	} else {
		fmt.Println("userID is an", val)
	}

}
