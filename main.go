// Interfaces are contracts that will implicitly garauntee any value is going to have a method
// Interface will go find the unique implmentation on the concrete type

package main

import "fmt"

// a concrete person type
type person struct {
	first string
}

// method: person speak
func (p person) speak() {
	fmt.Println("My Person Name", p.first)
}

// another concrete parent type
// with an abstract type of person
type parent struct {
	person
	saysNo bool
}

// method: parent speak
func (p parent) speak() {
	fmt.Println("My Parent Name", p.first)
}

//an interfce says
// any type that has the methods specified by an interface,
// ...is also of the interface type
// If you have these methods, then you are my type

// an abstract human type
type human interface {
	speak()
}

//now create a function
func foo(h human) {
	h.speak()
}

func main() {

	//a variable p1 contains a value of type person
	p1 := person{
		first: "James",
	}

	p2 := parent{
		person: person{
			first: "Dad",
		},
		saysNo: true,
	}

	//fmt.Printf("Here is a person type %T\n", p1)

	// a value can be of more than 1 type
	// in this example, p1 is both type person and type human
	var x human //x is a human
	x = p1      //is now a person too
	x.speak()   //call human speak, which calls person speak

	x = p2    //is now a parent too
	x.speak() //call human speak, which calls parent speak

	//substitutability...
	fmt.Println("Substitutability...")
	foo(x)  //explicit human interface
	foo(p1) //concrete person type
	foo(p2) //concrete parent type

	//Calling a concrete method explicity thru the abstract class
	// using .dot notation
	fmt.Println(".dot notation...")
	fmt.Println("here:", x.(parent).first)
	x.(parent).speak() //from human speak, calls person speak

}

//git status
//git add -A
//git commit - "added more interface"
//git push
//git tag or git log
//git tag v0.0.5
//git push --tags
