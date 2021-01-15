package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // this is the address of the memory where a is stored, & is the operator to see that

	fmt.Printf("%T\n", a)  // it returns int
	fmt.Printf("%T\n", &a) // it returns *int - "pointer to an int"

	var b *int = &a // here the * is part of the type
	fmt.Println(b)
	fmt.Println(*b) // here the * is to check what value is stored in this address, here the * is an operator - unreferencing an address
	fmt.Println(*&a)
	*b = 53
	fmt.Println(a)

	fmt.Println("When to use pointers")
	x := 0
	foo(&x)
	fmt.Println(x)

	fmt.Println("Methods set")
}

func foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}
