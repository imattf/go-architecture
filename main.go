// Stringer
// overriding the Println function for a struct type
// using the String() method
// see fmt.Stringer https://github.com/golang/go/blob/go1.14.4/src/fmt/print.go#L62

package main

import "fmt"

type person struct {
	first string
}

//If you have the String() method, Println will use this to print out values
func (p person) String() string {
	return fmt.Sprintf("My name is %s", p.first)
}

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	p1 := person{
		first: "Jenny",
	}

	fmt.Println(p1)

	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
	// Output: Gopher (2)

}
