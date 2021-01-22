package main

import "fmt"

func main() {
	fmt.Println("Channels:")

	//c := make(chan int)
	//go func(){
	//	c <- 42
	//}()

	// we put one on buffer
	c := make(chan int, 2)
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	// directional channels
	fmt.Println("---------------------")
	fmt.Printf("%T\n", c)

	// send only
	//c := make(chan <- int, 2)

	// receive only
	//c := make(<- chan int, 2)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// general to specific converts but not the other way around
	fmt.Println("---------------------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}
