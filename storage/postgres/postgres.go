package postgres

import (
	"github.com/ag3ntsc4rn/go-template"
)

type Db map[int]template.Person

func (pg Db) Save(p template.Person, n int) {
	pg[n] = p
}

func (pg Db) Retrieve(n int) template.Person {
	return pg[n]
}
