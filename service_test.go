package template

import(
	"testing"
)

type Db map[int] Person

func (m Db) Save(p Person, n int) {
	m[n] = p
}

func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	db := Db{}
	ps := NewPersonService(db)
	p := Person {
		First: "Michael",
	}
	ps.Put(p, 1)
	got, _ := ps.Get(1)
	if got != p {
		t.Fatalf("Got %v, Want %v", got, p)
	}
	
}