package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("Ex.07.01")
	x := "Hello"
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println("Ex.07.02")
	p1 := person{
		first: "Todd",
		last:  "Apple",
		age:   50,
	}
	fmt.Println(p1.first, p1.last, p1.age)
	changeMe(&p1)
	fmt.Println(p1.first, p1.last, p1.age)
	changeMe2(&p1)
	fmt.Println(p1.first, p1.last, p1.age)
}

func changeMe(p *person) {
	*p = person{
		first: "James",
		last:  "Bond",
		age:   35,
	}
}

func changeMe2(p *person) {
	//p.first = "Miss"
	(*p).first = "Miss Penny"
}
