// Hands-on exercise #4 - Ninja Level 3
// create a background context
// create a child context with a timeout
// show the code working two ways
// not timing out
// timing out
// Things that will be useful to use in this example:
// time.Time is a specific point in time
// time.Duration is a period of time
// time.Millisecond
// this is a duration
// time.Sleep
// takes in a duration
// tag v1.15.0 & v1.17.0
// video: 71 Hands On 4 - Attempt 2

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background() //default parent

	// this is the work
	ctx, cancelf := context.WithTimeout(ctx, 100*time.Second)
	defer cancelf()

	time.Sleep(5 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("work timed-out reached")
	default:
		fmt.Println("work finished")
	}
}
