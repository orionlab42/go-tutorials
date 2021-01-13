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

	// anonymous func
	func(x int) {
		fmt.Println("Anonymous func ran, meaning of life: ", x)
	}(42)

	// func expression
	f := func(x int) {
		fmt.Println("My first func expression", x)
	}
	f(1848)

	s1 := foo2()
	fmt.Println(s1)
	s2 := bar3(2000)
	i := s2()
	fmt.Printf("%T\n", s2)
	fmt.Printf("%v\n", bar3(2000)())
	fmt.Printf("%v\n", s2())
	fmt.Printf("%v\n", i)
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := sum1(ii...)
	fmt.Println(s4)
	s5 := even(sum1, ii...)
	fmt.Println(s5)
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	n := factorial(4)
	fmt.Println(n)
	fmt.Println(loopFact(4))
}

// returning a func
func foo2() string {
	s := "Hello world"
	return s
}

func bar3(x int) func() int {
	return func() int {
		return x
	}
}

// callback - when you pass a func as an argument
func sum1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total = total + v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

// closure -- limiting scope
// code block in code block if you put in a function {} - the variables declared there are only working there

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// recursion  -- the function calls itself
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	fact := 1
	for i := n; i > 0; i-- {
		fact *= i
	}
	return fact
}
