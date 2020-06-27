//postgres stuff...
package postgres

import "github.com/imattf/go-architecture"

type Db map[int]architecture.Person

func (pg Db) Save(n int, p architecture.Person) {
	pg[n] = p
}
func (pg Db) Get(n int) architecture.Person {
	return pg[n]
}
