package harddrive

import (
	"fmt"

	"github.com/imattf/go-architecture"
)

// Db is the data store type hard-drive
type Db map[int]architecture.User

// Save is the data store save method
func (h Db) Save(n int, u architecture.User) {
	fmt.Println("...saving in hard-drive:", u)
	h[n] = u
}

// Retrieve is the data store retrive method
func (h Db) Retrieve(n int) architecture.User {
	fmt.Println("...retrieving from hard-drive:", h[n])
	return h[n]
}
