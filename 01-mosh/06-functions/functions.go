package main

import "fmt"

func main() {
	fmt.Println("Functions-syntax")
	text := bar1("James")
	fmt.Println(text)
	a, b := mouse("Ian", "Flemming")
	fmt.Println(a)
	fmt.Println(b)
	sum(2, 3, 4, 7, 8, 9)
	xi := []int{2, 3, 4, 5, 6, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is ", x)
	defer foo()
	bar()
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func bar1(s string) string {
	return fmt.Sprint("Hello ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says: Hello")
	b := false
	return a, b
}

// variadic parameters - it creates a slice

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("for item in index position, ", i, " we are adding, ", v, " to the total which is now ", sum)
	}
	return sum
}

// defer - runs a function when the containing function exits
func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
