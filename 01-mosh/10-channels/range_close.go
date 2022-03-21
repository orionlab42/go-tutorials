package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo1(c)
	// receive
	// receive
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}

//send
func foo1(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
}
