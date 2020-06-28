package architecture

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test setup...
// Copy local database mock into here, otherwise we get a import error
type Db map[int]Person

func (mg Db) Save(n int, p Person) {
	mg[n] = p
}
func (mg Db) Get(n int) Person {
	return mg[n]
}

func TestSomething(t *testing.T) {
	// All these assertions pass
	assert.Equal(t, "hello", "hello", "Values are equal")
	// assert.NotEqual(t, "hello", "world", "Values are different")
	// assert.Contains(t, "hello", "el", "String contains other given string")
	// assert.True(t, true, "Value is true")
	// assert.False(t, false, "Value is false")

	// All these assertions fail
	// assert.Equal(t, "hello", "world", "Values are equal")
	// assert.NotEqual(t, "hello", "hello", "Values are different")
	// assert.Contains(t, "hello", "y", "String contains other given string")
	// assert.True(t, false, "Value is true")
	// assert.False(t, true, "Value is false")

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

	assert.True(t, true, "True is true!")

}

//this will create Example in go doc...
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

func TestGet(t *testing.T) {
	mdb := Db{}
	p := Person{
		First: "Bob",
	}
	Put(mdb, 1, p)

	got := mdb.Get(1)
	if got != p {
		t.Fatalf("Want %v, got %v", p, got)
	}

}

//this will create Example in go doc...
func ExampleGet() {
	mdb := Db{}
	p := Person{
		First: "Bob",
	}
	Put(mdb, 1, p)

	got := mdb.Get(1)
	fmt.Println(got)
	// Output: {Bob}

}
