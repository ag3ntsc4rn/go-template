package template

import (
	"fmt"
)

type Person struct {
	First string
}

type Accessor interface {
	Save(p Person, n int)
	Retrieve(n int) Person
}

func Put(a Accessor, p Person, n int) {
	a.Save(p, n)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService{
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Put(p Person, n int ){
	ps.a.Save(p, n)
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No person found at index %d", n)
	}
	return p, nil
}
