package main

import (
	"fmt"
	_2_program_structure "github.com/annakallo/go-tutorials/02-donovan-kernighan-book/02-program-structure"
)

type bum struct {
	Name string
}

func main() {
	x := _2_program_structure.Gcd(15, 9)
	fmt.Println(x)
}
