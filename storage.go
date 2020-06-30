package architecture

// User is the data store object type
type User struct {
	First string
}

// Accessor is how data store accesses the user info
type Accessor interface {
	Save(n int, u User)
	Retrieve(n int) User
}

// Put function takes an interface type accessor, stores a user value
func Put(a Accessor, n int, u User) {
	a.Save(n, u)
}

// Get function takes an interface type accessor, gets a user value
func Get(a Accessor, n int) User {
	return a.Retrieve(n)
}
