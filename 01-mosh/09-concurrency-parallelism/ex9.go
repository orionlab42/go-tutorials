package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type person struct {
	name     string
	eyeColor string
	weight   int
}

func (p *person) speak() {
	fmt.Println("Hello my name is ", p.name, " and my eye color is ", p.eyeColor)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

var wg2 sync.WaitGroup

func main() {
	fmt.Println("Ex.09.01")
	//fmt.Println("Goroutines: ", runtime.NumGoroutine())
	//wg2.Add(2)
	//go routine1()
	//fmt.Println("Goroutines: ", runtime.NumGoroutine())
	//go routine2()
	//fmt.Println("Goroutines: ", runtime.NumGoroutine())
	//wg2.Wait()
	//fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Ex.09.02")
	p1 := person{
		name:     "Todd",
		eyeColor: "green",
		weight:   80,
	}
	saySomething(&p1)
	// doesn't not work:
	//saySomething(p1)
	p1.speak()
	fmt.Println("Ex.09.03")
	var incrementer int64
	wg2.Add(100)
	//var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			//mu.Lock()
			atomic.AddInt64(&incrementer, 1)
			//plus := incrementer
			//runtime.Gosched() - commented for the mutex
			//plus++
			//incrementer = plus
			//fmt.Println("Incrementer mutex: ex 09.04 ", incrementer)
			fmt.Println("Incrementer atomic: ex 09.05 ", atomic.LoadInt64(&incrementer))
			//mu.Unlock()
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println(incrementer)
	fmt.Println("Ex.09.04")
	fmt.Println("Ex.09.05")
	fmt.Println("Ex.09.06")
	fmt.Println("Ex.09.07")
}

func routine1() {
	fmt.Println("AAAAAAAAAA: Routine 1")
	wg2.Done()
}

func routine2() {
	fmt.Println("BBBBBBBBBB: Routine 2")
	wg2.Done()
}
