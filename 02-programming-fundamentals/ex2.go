package main

import "fmt"

const (
	one int = 3
	two     = 42
)
const (
	yearOne = 2017 + iota
	yearTwo
	yearThree
	yearFour
)

func main() {
	fmt.Println(one, two)
	var a = 478
	fmt.Printf("%d\t\t%b\t\t%x\n", a, a, a)
	aShifted := a << 1
	fmt.Printf("%d\t\t%b\t\t%x\n", aShifted, aShifted, aShifted)
	bShifted := a << 2
	fmt.Printf("%d\t\t%b\t\t%x\n", bShifted, bShifted, bShifted)
	//fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)
	//b := 42 <= 42
	//c := 42 < 42
	//d := 42 != 42
	//fmt.Println(b, c, d)
	rawString := `"Raw string 
						wtf"`
	fmt.Println(rawString)
	fmt.Println(yearOne, yearTwo, yearThree, yearFour)
}
