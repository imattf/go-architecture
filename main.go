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

func store(d datastore, n int, p person) {
	d.save(n, p)
}
func get(d datastore, n int) person {
	return d.get(n)
}

func main() {

	//create the persons data
	fmt.Println("Creating the data...")
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

	//create the datastores
	fmt.Println("Create datastore...")
	mds := mongodb{}
	pds := postgres{}

	//save the persons in a mongodb datastore
	fmt.Println(".start storing data into mongodb datastore")
	store(mds, 1, p1)
	store(mds, 2, p2)
	store(mds, 3, p3)
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

	//store the persons in a postgres datastore
	fmt.Println(".start storing data into postres datastore")
	store(pds, 1, p1)
	store(pds, 2, p2)
	store(pds, 3, p3)
	fmt.Println("...storing:", p1)
	fmt.Println("...storing:", p2)
	fmt.Println("...storing:", p3)

	//get the persons from the postgres datastore
	fmt.Println(".start getting data from postgres datastore")
	rec1 = get(pds, 1)
	rec2 = get(pds, 2)
	rec3 = get(pds, 3)
	fmt.Println("...getting:", rec1)
	fmt.Println("...getting:", rec2)
	fmt.Println("...getting:", rec3)

}
