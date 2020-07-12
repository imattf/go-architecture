// Hands-on exercise #1
// create a type
// have it implement the error interface
// create a value of that type
// pass that into a function, or return it from a function, that takes type error
// tag v1.30.0
// video: 90 Hands On 1

package main

import (
	"fmt"
	"log"
)

type errorString string

func (es errorString) Error() string {
	return fmt.Sprintf("this is an error string who,what,when,where - %s", string(es))
}

func main() {
	n, err := sum(2, 4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func sum(i, j int) (int, error) {
	n := i * j
	if n != i+j {
		var sErr errorString = "numbers don't add up"
		return 0, sErr
	}
	return n, nil
}
