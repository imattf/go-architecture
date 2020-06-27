package architecture

import "fmt"

// Person is how the architecture package stores a person.
type Person struct {
	First string
}

// Datastore is how to store and get a person.
// When getting a person, if they do not exist, return the zero value.
type Datastore interface {
	Save(n int, p Person)
	Get(n int) Person
}

// DatastoreService is wrapper service to get records from a Datastore.
type DatastoreService struct {
	d Datastore
}

// Get method takes a key value and returns a record value.
// If no key exists, an error is returned.
func (ds DatastoreService) Get(n int) (Person, error) {
	p := ds.d.Get(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no Person n %d", n)
	}
	return p, nil
}

// NewDatastoreService function takes a Datastore and returns a new DatastoreService.
func NewDatastoreService(d Datastore) DatastoreService {
	return DatastoreService{
		d: d,
	}
}

// Get function takes a Datastore and key value and returns a record value.
func Get(d Datastore, n int) Person {
	return d.Get(n)
}

// Put function takes a Datastore, key value and a record value.
func Put(d Datastore, n int, p Person) {
	d.Save(n, p)
}
