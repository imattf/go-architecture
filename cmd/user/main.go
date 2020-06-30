package main

import (
	"fmt"

	"github.com/imattf/go-architecture"
	"github.com/imattf/go-architecture/storage/mongo"
	// "github.com/imattf/go-architecture/storage/harddrive"
)

func main() {

	// create the data store type...
	storage := mongo.Db{} //use a mongo mock
	// storage := harddrive.Db{} //use a hard-drive mock

	// create values to store in data store...
	u1 := architecture.User{
		First: "James",
	}

	u2 := architecture.User{
		First: "Jenny",
	}

	// put the values in the data store..
	fmt.Println("Putting values into data storage...")
	architecture.Put(storage, 1, u1)
	architecture.Put(storage, 2, u2)

	// get the values in the data store..
	fmt.Println("Getting values from data storage...")
	// fmt.Println(get(storage, 1))
	// fmt.Println(get(storage, 2))
	architecture.Get(storage, 1)
	architecture.Get(storage, 2)

}
