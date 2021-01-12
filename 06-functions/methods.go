package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

// interfaces
type human interface {
	speak()
}

func bar2(h human) {
	fmt.Println("I was passed into bar2 ", h)
	// assertion
	switch h.(type) {
	case person:
		fmt.Println("Person ", h.(person).first)
	case secretAgent:
		fmt.Println("Secret agent ", h.(secretAgent).first)
	}
}

// in an empty interface you can put every value, because every type has no methods

// func (r receiver) identifier(parameters) (return(s)) {code}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func (s person) speak() {
	fmt.Println("I am ", s.first, s.last, " person speak")
}

func main() {
	fmt.Println("Methods:")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	p1 := person{
		first: "Dr",
		last:  "Yes",
	}
	fmt.Println(p1)
	p1.speak()

	// polymorphism
	bar2(sa1)
	bar2(p1)

}
