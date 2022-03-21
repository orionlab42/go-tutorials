package main

import (
	"fmt"
	"math"
	"strconv"
)

type persons struct {
	first string
	last  string
	age   int
}
type square struct {
	sideLength float64
}
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {
	fmt.Println("Ex.06.01")
	//fooo1 := fooo()
	//fmt.Println("Returned by fooo: ", fooo1)
	//baaar1, baaar2 := baaar()
	//fmt.Println("Returned by baaar: ", baaar1, baaar2)
	fmt.Println("Ex.06.02")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fooo(xi...)
	//defer fooo(xi...)
	baaar(xi)
	fmt.Println("Ex.06.03")
	p1 := persons{
		first: "Todd",
		last:  "Something",
		age:   53,
	}
	p1.speak()
	fmt.Println("Ex.06.04")
	sq := square{
		sideLength: 5,
	}
	circ := circle{
		radius: 3,
	}
	info("Square area: ", sq)
	info("Circle area: ", circ)
	fmt.Println("Ex.06.05")
	func() {
		fmt.Println("Hello anonymus!!!")
	}()
	fmt.Println("Ex.06.06")
	f := func(s int) int {
		sumOdd := 0
		for i := 0; i <= s; i++ {
			if i%2 == 1 {
				sumOdd += i
			}
		}
		return sumOdd
	}
	fmt.Println(f(100))
	fmt.Printf("%T\n", f)
	fmt.Println("Ex.06.07")
	funkyBiz := ex8(8)
	funkyBiz("Unicorn")
	fmt.Println(funkyBiz("Unicorn"))
	fmt.Println("Ex.06.08")
	xs := ex9(age, 1989)
	fmt.Println(xs)
	fmt.Println("Ex.06.09")
	aa := closureF()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	bb := closureF()
	fmt.Println(bb())
	fmt.Println(aa())
	fmt.Println("Ex.06.10")
}

//fmt.Println("Ex.06.01")

//func fooo() int {
//	return 200
//}
//func baaar() (int, string) {
//	return 400, "Yesss"
//}

//fmt.Println("Ex.06.02")

func fooo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("Sum fooo= ", sum)
	return sum
}

func baaar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("Sum baaar= ", sum)
	return sum
}

//fmt.Println("Ex.06.03")

func (p persons) speak() {
	fmt.Println("My name is ", p.first, p.last, " and I have ", p.age, " years.")
}

//fmt.Println("Ex.06.04")

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (s circle) area() float64 {
	return s.radius * s.radius * math.Pi
}

func info(s string, sh shape) {
	fmt.Println(s, sh.area())
}

//fmt.Println("Ex.06.05")
//fmt.Println("Ex.06.06")
//fmt.Println("Ex.06.07")

func ex8(n int) func(s string) string {
	return func(s string) string {
		fmt.Println("You choose the number: ", n)
		mes := "Better stop at this word: " + s
		return mes
	}
}

//fmt.Println("Ex.06.08")
func ex9(g func(n int) int, yB int) string {
	years := g(yB)
	age := "You are " + strconv.Itoa(years) + " old."
	return age
}

func age(yearBirth int) int {
	age := 2020 - yearBirth
	return age
}

//fmt.Println("Ex.06.09")
func closureF() func() int {
	var x int
	return func() int {
		x += 1
		x = x * x
		return x
	}
}

//fmt.Println("Ex.06.10")
