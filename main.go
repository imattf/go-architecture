// Hands-on exercise #5 - Ninja Level 3
// create a background context
// create a child context with cancel
// assign this to a variable “ctx”
// launch 100 goroutines
// give a unique number to each goroutine
// print “running” and the number
// have an eternal loop with
// select
// case <- ctx.Done():
// return
// default:
// print “still working” and the number
// sleep for a 1/10 of a second
// show the number of goroutines running
// after your eternal for loop
// sleep for a bit so that all goroutines can be launched
// call your cancel func
// this will cancel your context, and all launched goroutines
// sleep for a bit so that your program can close goroutines
// show the goroutines running
// exit main

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Printf("Number of goroutines running: %d\n", runtime.NumGoroutine())

	ctx := context.Background() //default parent
	ctx, cancelf := context.WithCancel(ctx)
	defer cancelf()

	for i := 0; i < 100; i++ {
		//anonymous func
		go func(n int) {
			fmt.Println("launching goroutine", n)
			for {
				select {
				case <-ctx.Done():
					return
					//runtime.Goexit()
				default:
					fmt.Printf("...goroutine %d is doing work\n", n)
					time.Sleep(2 * time.Second)
				}
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Printf("Number of goroutines running after 1 Millisecond: %d\n", runtime.NumGoroutine())
	cancelf()
	fmt.Println("Cancelling goroutines")
	time.Sleep(5 * time.Second)
	fmt.Printf("Number of goroutines running after cancelf() called: %d\n", runtime.NumGoroutine())
}
