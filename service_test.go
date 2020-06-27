package architecture

import (
	"fmt"
	"testing"
)

// Copy local database mock into here, otherwise we get a import error

type Db map[int]Person

func (mg Db) Save(n int, p Person) {
	mg[n] = p
}
func (mg Db) Get(n int) Person {
	return mg[n]
}

func TestPut(t *testing.T) {
	mdb := Db{}
	p := Person{
		First: "Bob",
	}

	Put(mdb, 1, p)
	got := mdb.Get(1)
	if mdb.Get(1) != p {
		t.Fatalf("Want %v, got %v", p, got)
	}

}

func ExamplePut() {
	mdb := Db{}
	p := Person{
		First: "Bob",
	}

	Put(mdb, 1, p)
	got := mdb.Get(1)
	fmt.Println(got)
	// Output: {Bob}

}
