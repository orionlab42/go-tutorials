package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Fan in: ")
	// Collecting many channels into one

	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send2(eve, odd)

	//receive
	go receive2(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func send2(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive2(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
