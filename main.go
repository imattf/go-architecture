//Parent Context context.Background()
// with a child context given a Value
// a better approach

package main

import (
	"context"
	"fmt"

	"github.com/imattf/go-architecture/session"
)

type key int

var userKey key = 0

func main() {

	ctx := context.Background()       //no Value associated w/ context, so you get "not logged in"
	ctx = session.WithUserID(ctx, 42) //you will be logged in

	userID := session.GetUserID(ctx)
	if userID == nil {
		fmt.Println("Not logged in!")
		return
	}
	fmt.Println(*userID)

	uID := ctx.Value(userKey)
	fmt.Println(uID)

}
