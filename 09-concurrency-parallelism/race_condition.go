package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Race condition")
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//v := counter
			//time.Sleep(time.Second*2)
			runtime.Gosched()
			//v++
			//counter = v
			counter++
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	fmt.Println("CPUs", runtime.NumCPU())

}
