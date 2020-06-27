// Interfaces are contracts that will implicitly garauntee any value is going to have a method
// Interface will go find the unique implmentation on the concrete type
// datastore example

package main

import "fmt"

//datastore schema
type person struct {
	first string
}

//mongodb stuff...
type mongodb map[int]person

func (mg mongodb) save(n int, p person) {
	mg[n] = p
}
func (mg mongodb) get(n int) person {
	return mg[n]
}

//postgres stuff...
type postgres map[int]person

func (pg postgres) save(n int, p person) {
	pg[n] = p
}
func (pg postgres) get(n int) person {
	return pg[n]
}

//datastore stuff...
type datastore interface {
	save(n int, p person) //pass in key & person value
	get(n int) person     //pass in key & return a person value
}

func put(d datastore, n int, p person) {
	d.save(n, p)
}
func get(d datastore, n int) person {
	return d.get(n)
}

func main() {

	//create the datastore
	fmt.Println("Create datastore...")
	mds := mongodb{}
	//pds := postgres{}

	//create the persons
	fmt.Println(".start creating the data")
	p1 := person{
		first: "Bob",
	}
	p2 := person{
		first: "Terry",
	}
	p3 := person{
		first: "Cathy",
	}
	fmt.Println("...created:", p1)
	fmt.Println("...created:", p2)
	fmt.Println("...created:", p3)

	//save the persons in a mongodb datastore
	fmt.Println(".start loading data into mongodb datastore")
	put(mds, 1, p1)
	put(mds, 2, p2)
	put(mds, 3, p3)
	fmt.Println("...storing:", p1)
	fmt.Println("...storing:", p2)
	fmt.Println("...storing:", p3)

	//get the persons from the mongodb datastore
	fmt.Println(".start getting data from mongodb datastore")
	rec1 := get(mds, 1)
	rec2 := get(mds, 2)
	rec3 := get(mds, 3)
	fmt.Println("...getting:", rec1)
	fmt.Println("...getting:", rec2)
	fmt.Println("...getting:", rec3)

}
