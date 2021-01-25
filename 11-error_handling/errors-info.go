package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMath = errors.New("error as a variable")

func main() {
	fmt.Printf("%T\n", ErrMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		//return 0, errors.New("square root of negative number")
		//return 0, ErrMath
		return 0, fmt.Errorf("error as a fmt: %v", n)
	}
	return 42, nil
}
