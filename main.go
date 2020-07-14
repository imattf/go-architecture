// Stack, Heap and Escape Analysis
// with go run -gcflags "-m" main.go

package main

//go:noinline

func foo() int {
	x := 42
	return x
}

//go:noinline
func bar() *int {
	y := 21
	return &y
}

func main() {

	//prevent Go complier from optimizing the code here
	_ = foo() //on return gets removed from stack
	_ = bar() //on return y gets copied to heap

}

// result:
// go-architecture % go run -gcflags "-m" main.go
// # command-line-arguments
// ./main.go:14:2: moved to heap: y
// go-architecture %
