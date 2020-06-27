// Interfaces are contracts that will implicitly garauntee any value is going to have a method
// Interface will go find the unique implmentation on the concrete type
// Datastore example w/ refactor for DatastoreService with a get method...

package architecture

import "fmt"

//Datastore schema
type Person struct {
	First string
}

//Datastore stuff...
type Datastore interface {
	Save(n int, p Person) //pass in key & Person value
	Get(n int) Person     //pass in key & return a Person value
}

//wrap service around Datastore...
type DatastoreService struct {
	d Datastore
}

func NewDatastoreService(d Datastore) DatastoreService {
	return DatastoreService{
		d: d,
	}
}

//... with a get service
func (ds DatastoreService) Get(n int) (Person, error) {
	p := ds.d.Get(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no Person n %d", n)
	}
	return p, nil
}

func Store(d Datastore, n int, p Person) {
	d.Save(n, p)
}
func Get(d Datastore, n int) Person {
	return d.Get(n)
}
