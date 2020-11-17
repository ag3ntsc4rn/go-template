package mongo

import (
	"github.com/ag3ntsc4rn/go-template"
)

type Db map[int]template.Person

func (m Db) Save(p template.Person, n int) {
	m[n] = p
}

func (m Db) Retrieve(n int) template.Person {
	return m[n]
}