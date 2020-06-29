// Hand-on Exercise 2 - Ninja Level 2
// Recreate the database example using two maps to simulate the databases. 
// Create an “accessor” interface that has two methods
//  save
//  retrieve
// have your databases implement the “accessor” interface
// create functions that take in values of type accessor
//  put
//  get
// use “put” and “get” to store data in either database


package main

import (
	"fmt"
)

// data store object
type user struct {
	first string
}

// data store type mongo
type mongo map[int]user

// data store save method
func (m mongo) save(n int, u user) {
	m[n] = u
	fmt.Println("...saving in mongo:", m[n])
}

// data store retrive method
func (m mongo) retrieve(n int) user {
	fmt.Println("...retrieving from mongo:", m[n])
	return m[n]
}

// data store type hard-drive
type hdrive map[int]user

// data store save method
func (h hdrive) save(n int, u user) {
	fmt.Println("...saving in hard-drive:", u)
	h[n] = u
}

// data store retrive method
func (h hdrive) retrieve(n int) user {
	fmt.Println("...retrieving from hard-drive:", h[n])
	return h[n]
}

// data store interface type
type accessor interface {
	save(n int, u user)
	retrieve(n int) user
}

// Put function takes an interface type accessor, stores a value 
func put(a accessor, n int, u user) {
	a.save(n, u)
}

// Get function takes an interface type accessor, gets a value
func get(a accessor, n int) user {
	return a.retrieve(n)
}

func main() {

	// create the data store type...
	storage := mongo{}  //use a mongo mock
	// storage := hdrive{} //use a hard-drive mock

	// create values to store in data store...
	u1 := user{
		first: "James",
	}

	u2 := user{
		first: "Jenny",
	}

	// put the values in the data store..
	fmt.Println("Putting values into data storage...")
	put(storage, 1, u1)
	put(storage, 2, u2)

	// get the values in the data store..
	fmt.Println("Getting values from data storage...")
	// fmt.Println(get(storage, 1))
	// fmt.Println(get(storage, 2))
	get(storage, 1)
	get(storage, 2)

}
