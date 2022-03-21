package main

import "fmt"

func main() {
	fmt.Println("Ex.10.01")
	//c := make(chan int)
	//
	//go func(){
	//	c <- 42
	//}()
	//
	//fmt.Println(<-c)
	fmt.Println("Ex.10.02")
	//cs := make(chan <- int)
	//cr := make(<-chan  int)
	//go func(){
	//	cr <- 42
	//	cs <- 43
	//}()
	//fmt.Println(<-cr)
	//fmt.Println(<-cs)
	fmt.Println("Ex.10.03")
	//c := gen()
	//receive4(c)
	//fmt.Println("About to exit..")
	fmt.Println("Ex.10.04")
	//q := make(chan int)
	//c := gen1(q)
	//receive5(c, q)
	//fmt.Println("About to exit..")
	fmt.Println("Ex.10.05")
	//c := make(chan int)
	//go func(){
	//	c <- 42
	//}()
	//v, ok := <- c
	//fmt.Println(v, ok)
	//close(c)
	//v, ok = <- c
	//fmt.Println(v, ok)
	fmt.Println("Ex.10.06")
	//c := make(chan int)
	//go func(){
	//	for i:=0; i<100; i++{
	//		c <- i
	//	}
	//	close(c)
	//}()
	////for v := range c {
	////	fmt.Println(v)
	////}
	//for i:=0;i<10;i++ {
	//	fmt.Println(<-c)
	//}
	//fmt.Println("About to exit..")
	fmt.Println("Ex.10.07")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for k := 0; k < 10; k++ {
				c <- k
			}
		}()
	}

	for j := 0; j < 100; j++ {
		fmt.Println(<-c)
	}
	fmt.Println("About to exit..")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive4(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen1(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		//close(c)
		q <- 0
	}()
	return c
}

func receive5(c, q <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println("Number: ", v)
		case v := <-q:
			fmt.Println("From the quit channel: ", v)
			return
		}
	}
}
