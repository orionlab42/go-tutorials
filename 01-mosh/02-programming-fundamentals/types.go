package main

import "fmt"

const (
	a = iota
	b
	c = iota
)

const (
	e = iota
	d
	f
)

func main() {

	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)
	//
	//s := "Hello, playground"
	//fmt.Println(s)
	//fmt.Printf("%T\n", s)
	//
	//bs := []byte(s)
	//fmt.Println(bs)
	//fmt.Printf("%T\n", bs)
	//
	//for i := 0; i < len(s); i ++ {
	//	fmt.Printf("%#U", s[i])
	//}
	//fmt.Println("")
	//for i, v := range s {
	//	fmt.Printf("at index position %d we have %#x\n", i, v)
	//}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(d)
	fmt.Println(f)

}
