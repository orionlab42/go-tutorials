package main

import (
	"fmt"
	ch2_program_structure "github.com/annakallo/go-tutorials/02-donovan-kernighan-book/ch2-program-structure"
)

type bum struct {
	Name string
}

func main() {
	x := ch2_program_structure.CToF(160)
	y := ch2_program_structure.FToC(320)
	fmt.Println(x)
	fmt.Println(ch2_program_structure.Fahrenheit(y))
	fmt.Println(x > ch2_program_structure.Fahrenheit(y))
}
