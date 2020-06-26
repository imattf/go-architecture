package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

//an interfce says
// any type that has the methods specified by an interface,
// ...is also of the interface type
// If you have these methods, then you are my type
type human interface {
	speak()
}

func main() {

	//a variable p1 contains a value of type person
	p1 := person{
		first: "James",
	}
	fmt.Printf("Here is a person type %T\n", p1)

	// a value can be of more than 1 type
	// in this example, p1 is both type person and type human
	var x human
	x = p1
	fmt.Printf("Here is a human type %T\n", x)
}
