package main

import "fmt"

func main() {
	//for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("outer loop %v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t inner loop %v\t  %v\n", i, j)
		}
	}
	x := 0
	for x <= 122 {
		fmt.Printf("Unicode for %v is %#U and in hexadecimal %#x\n", x, x, x)
		x++
	}

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print2")
	case (3 == 3):
		fmt.Println("This prints")
		fallthrough
	case (4 == 5):
		fmt.Println("This prints2")
		fallthrough
	default:
		fmt.Println("This is default")
	}

}
