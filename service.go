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
