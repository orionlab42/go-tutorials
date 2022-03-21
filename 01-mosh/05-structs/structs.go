package main

import "fmt"

type person struct {
	first string
	last  string
}

// embedded structs
type secretAgent struct {
	person
	ltk bool
}

func main() {
	fmt.Println("Structs")
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Penny",
	}
	//fmt.Println(p1, p2)
	fmt.Println(p1.first)
	fmt.Println(p2.first)
	//fmt.Printf("%T\n", p1)

	sa1 := secretAgent{
		person: person{
			first: "Todd",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	// inner type gets promoted to the outer type --> you dont need to specify that first is in person, only if there is a first also in secretAgent
	fmt.Println(sa1.first, sa1.last, sa1.ltk)

	// anonymous structs -- it's good when you dont want to pollute your code
	t1 := struct {
		name string
		age  int
	}{
		name: "Peter",
		age:  55,
	}
	fmt.Println(t1)
	fmt.Println(t1.name, t1.age)

}
