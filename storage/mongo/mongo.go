//mongodb stuff...
package mongo

import "github.com/imattf/go-architecture"

type Db map[int]architecture.Person

func (mg Db) Save(n int, p architecture.Person) {
	mg[n] = p
}
func (mg Db) Get(n int) architecture.Person {
	return mg[n]
}
