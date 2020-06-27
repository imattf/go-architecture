// Interfaces are contracts that will implicitly garauntee any value is going to have a method
// Interface will go find the unique implmentation on the concrete type
// datastore example w/ refactor for datastoreService with a get method...

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

//wrap service around datastore...
type datastoreService struct {
	d datastore
}

//... with a get service
func (ds datastoreService) get(n int) (person, error) {
	p := ds.d.get(n)
	if p.first == "" {
		return person{}, fmt.Errorf("no person n %d", n)
	}
	return p, nil
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

	//create the mongodb datastore
	fmt.Println("Create mongodb datastore...")
	mds := mongodb{}

	//...now create a get service for the mongodb datastore
	ds := datastoreService{
		d: mds,
	}

	//save the persons in a mongodb datastore
	fmt.Println(".start storing data into mongodb datastore")
	store(mds, 1, p1)
	store(mds, 2, p2)
	store(mds, 3, p3)
	fmt.Println("...storing 1:", p1)
	fmt.Println("...storing 2:", p2)
	fmt.Println("...storing 3:", p3)

	//get the persons from the mongodb datastore
	fmt.Println(".start getting data from mongodb datastore")
	// rec1 := get(mds, 1)
	// rec2 := get(mds, 2)
	// rec3 := get(mds, 3)
	fmt.Println("...getting 1:", get(mds, 1))
	fmt.Println("...getting 2:", get(mds, 2))
	fmt.Println("...getting 3:", get(mds, 3))

	//TODO: refactor this stuff so Println can prints a string message
	//now use the datastoreService get
	// fmt.Println(".start getting data from datastoreService")
	// rec1, err := ds.get(1)
	// if err != nil {
	// 	fmt.Println("...getting 1:", err)
	// } else {
	// 	fmt.Println("...getting 1:", rec1)
	// }
	//...

	//...now use the mongodb datastoreService get
	fmt.Println(".start getting data from mongodb datastoreService")
	fmt.Println(ds.get(1))
	fmt.Println(ds.get(2))
	fmt.Println(ds.get(3))
	fmt.Println(ds.get(4)) //show exception

	//create the mongodb datastore
	fmt.Println("Create postgres datastore...")
	pds := postgres{}

	//...now create a get service for the postgres datastore
	ds = datastoreService{
		d: pds,
	}

	//store the persons in a postgres datastore
	fmt.Println(".start storing data into postres datastore")
	store(pds, 1, p1)
	store(pds, 2, p2)
	store(pds, 3, p3)
	fmt.Println("...storing 1:", p1)
	fmt.Println("...storing 2:", p2)
	fmt.Println("...storing 3:", p3)

	//get the persons from the postgres datastore
	fmt.Println(".start getting data from postgres datastore")
	// rec1 = get(pds, 1)
	// rec2 = get(pds, 2)
	// rec3 = get(pds, 3)
	fmt.Println("...getting: 1", get(pds, 1))
	fmt.Println("...getting: 2", get(pds, 2))
	fmt.Println("...getting: 3", get(pds, 3))

	//...now use the datastoreService get
	fmt.Println(".start getting data from postgres datastoreService")
	fmt.Println(ds.get(1))
	fmt.Println(ds.get(2))
	fmt.Println(ds.get(3))
	fmt.Println(ds.get(4)) //show exception

}
