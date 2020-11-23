package main

import "fmt"

type mytype int

var y int

func main() {

	var m mytype
	fmt.Printf("%v\t%T\t", m, m)
	m = 42
	fmt.Printf("%v\t%T\t", m, m)

	y = int(m)
	fmt.Printf("%v\t%T", y, y)
}
