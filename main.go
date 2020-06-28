// Hand on Exercise 1 - Ninja Level 2
// Create an interface and use it

package main

import (
	"fmt"
)

type animal struct {
	name  string
	walks bool
}

type dog struct {
	animal
	digs bool
}

func (dg dog) speak() {
	fmt.Println(dg.name, "says bark...")
	fmt.Println("...and do I like walks?, ", dg.walks)
	fmt.Println("...and do I dig in the yard?, ", dg.digs)
}
func (dg dog) showAffection() {
	fmt.Println("...and I show affection by wagging my tail")
}

type cat struct {
	animal
	hairball bool
}

func (ct cat) speak() {
	fmt.Println(ct.name, "says meow...")
	fmt.Println("...and do I like walks?, ", ct.walks)
	fmt.Println("...and do I pukes up hairballs?, ", ct.hairball)
}
func (ct cat) showAffection() {
	fmt.Println("...and I show affection by purring")
}

type pet interface {
	speak()
	showAffection()
}

func petAnalysis(p pet) {
	p.speak()
	p.showAffection()
}

func main() {
	pet1 := dog{
		animal: animal{
			name:  "Spot",
			walks: true,
		},
		digs: true,
	}

	pet2 := cat{
		animal: animal{
			name:  "Tiger",
			walks: false,
		},
		hairball: true,
	}

	fmt.Println("dog:", pet1)
	fmt.Println("cat:", pet2)

	petAnalysis(pet1)
	petAnalysis(pet2)

}
