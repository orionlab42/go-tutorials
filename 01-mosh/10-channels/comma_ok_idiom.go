package main

import "fmt"

func main() {
	fmt.Println("Comma ok idiom: / title")
	// First simple example
	//c := make(chan int)
	//go func(){
	//	c <- 42
	//	close(c)
	//}()
	//v, ok := <- c
	//fmt.Println(v, ok)
	//v, ok = <- c
	//fmt.Println(v, ok)

	// Second example
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send1(eve, odd, quit)

	//receive
	receive1(eve, odd, quit)

	fmt.Println("About to exit")
}

func send1(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive1(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel: ", v)
		case v := <-o:
			fmt.Println("From the odd channel: ", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok: ", i, ok)
				return
			} else {
				fmt.Println("From comma ok: ", i, ok)
			}
		}
	}
}
