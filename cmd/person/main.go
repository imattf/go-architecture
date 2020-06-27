package main

import (
	"fmt"

	"github.com/imattf/go-architecture"
	"github.com/imattf/go-architecture/storage/mongo"
	"github.com/imattf/go-architecture/storage/postgres"
)

func main() {

	//create the Persons data
	fmt.Println("Creating the data...")
	p1 := architecture.Person{
		First: "Bob",
	}
	p2 := architecture.Person{
		First: "Terry",
	}
	p3 := architecture.Person{
		First: "Cathy",
	}
	fmt.Println("...created:", p1)
	fmt.Println("...created:", p2)
	fmt.Println("...created:", p3)

	fmt.Println(" ")

	//create the mongodb Datastore
	fmt.Println("Create mongodb Datastore...")
	mds := mongo.Db{}

	//...now create a get service for the mongodb Datastore
	ds := architecture.NewDatastoreService(mds)

	//save the Persons in a mongodb Datastore
	fmt.Println(".start storing data into mongodb Datastore")
	architecture.Store(mds, 1, p1)
	architecture.Store(mds, 2, p2)
	architecture.Store(mds, 3, p3)
	fmt.Println("...storing 1:", p1)
	fmt.Println("...storing 2:", p2)
	fmt.Println("...storing 3:", p3)

	//get the Persons from the mongodb Datastore
	fmt.Println(".start getting data from mongodb Datastore")
	// rec1 := get(mds, 1)
	// rec2 := get(mds, 2)
	// rec3 := get(mds, 3)
	fmt.Println("...getting 1:", architecture.Get(mds, 1))
	fmt.Println("...getting 2:", architecture.Get(mds, 2))
	fmt.Println("...getting 3:", architecture.Get(mds, 3))

	//...now use the mongodb DatastoreService get
	fmt.Println(".start getting data from mongodb DatastoreService")
	fmt.Println(ds.Get(1))
	fmt.Println(ds.Get(2))
	fmt.Println(ds.Get(3))
	fmt.Println(ds.Get(4)) //show exception

	fmt.Println(" ")

	//create the mongodb Datastore
	fmt.Println("Create postgres Datastore...")
	pds := postgres.Db{}

	//...now create a get service for the postgres Datastore
	ds = architecture.NewDatastoreService(pds)

	//store the Persons in a postgres Datastore
	fmt.Println(".start storing data into postres Datastore")
	architecture.Store(pds, 1, p1)
	architecture.Store(pds, 2, p2)
	architecture.Store(pds, 3, p3)
	fmt.Println("...storing 1:", p1)
	fmt.Println("...storing 2:", p2)
	fmt.Println("...storing 3:", p3)

	//get the Persons from the postgres Datastore
	fmt.Println(".start getting data from postgres Datastore")
	// rec1 = get(pds, 1)
	// rec2 = get(pds, 2)
	// rec3 = get(pds, 3)
	fmt.Println("...getting: 1", architecture.Get(pds, 1))
	fmt.Println("...getting: 2", architecture.Get(pds, 2))
	fmt.Println("...getting: 3", architecture.Get(pds, 3))

	//...now use the DatastoreService get
	fmt.Println(".start getting data from postgres DatastoreService")
	fmt.Println(ds.Get(1))
	fmt.Println(ds.Get(2))
	fmt.Println(ds.Get(3))
	fmt.Println(ds.Get(4)) //show exception

}
