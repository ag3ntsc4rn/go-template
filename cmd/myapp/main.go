package main

import (
	"fmt"
	"github.com/ag3ntsc4rn/go-template"
	_"github.com/ag3ntsc4rn/go-template/storage/mongo"
	"github.com/ag3ntsc4rn/go-template/storage/postgres"
)

func main(){
	p1 := template.Person{
		First: "James",
	}
	p2 := template.Person{
		First: "Jenny",
	}

	db := postgres.Db{}

	ps := template.NewPersonService(db)
	ps.Put(p1, 1)
	ps.Put(p2, 2)
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(2))

}