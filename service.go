package architecture

import "fmt"

// Person is how the architecture package stores a person.
type Person struct {
	First string
}

// Datastore is how to store and get a person.
// When getting a person, if they do not exist, return the zero value.
type Datastore interface {
	Save(n int, p Person) //pass in key & Person value
	Get(n int) Person     //pass in key & return a Person value
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

// Store function takes a datastore, key value and a record value.
func Store(d Datastore, n int, p Person) {
	d.Save(n, p)
}

// Get function takes a datastore and key value and returns a record value.
func Get(d Datastore, n int) Person {
	return d.Get(n)
}
