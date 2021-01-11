package main

import "fmt"

func main() {
	fmt.Println("Functions-syntax")
	text := bar("James")
	fmt.Println(text)
	a, b := mouse("Ian", "Flemming")
	fmt.Println(a)
	fmt.Println(b)
	foo(2, 3, 4, 7, 8, 9)
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func bar(s string) string {
	return fmt.Sprint("Hello ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says: Hello")
	b := false
	return a, b
}

// variadic parameters

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("for item in index position, ", i, " we are adding, ", v, " to the total which is now ", sum)
	}
	fmt.Println("total total ", sum)
	return sum
}
