//Hands-on Exercise #2 - Ninja Level 3
// Using the code from the previous exercise
// have each key be its own type
// abstract all of the context code into its own
// functions as applicable
// package, separate from main
// pull the values out
// tag v1.14.0 & v1.16.0
// video 69 Hands On 2 - Attempt 1

package main

import (
	"context"
	"fmt"

	"github.com/imattf/go-architecture/session"
)

func main() {

	ctx := context.Background() //default parent

	ctx = session.SetUserID(ctx, 42)      //go add value with child
	ctx = session.SetUserAdmin(ctx, true) //go add value with child
	i := session.GetUserID(ctx)           //go get context value
	b := session.GetUserAdmin(ctx)        //go get context value
	fmt.Printf("user %d is an admin: %t\n", i, b)

}
