package mongo

import (
	"fmt"

	"github.com/imattf/go-architecture"
)

// Db is the data store type mongo
type Db map[int]architecture.User

// Save is the data store save method
func (m Db) Save(n int, u architecture.User) {
	fmt.Println("...saving in mongo:", u)
	m[n] = u
}

// Retrieve is the data store retrive method
func (m Db) Retrieve(n int) architecture.User {
	fmt.Println("...retrieving from mongo:", m[n])
	return m[n]
}
