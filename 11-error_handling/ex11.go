package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	Last  string
	First string
}

type customError struct {
	Number  int
	MyError error
}

func (e customError) Error() string {
	return fmt.Sprintf("a new error occured %v, %v:", e.Number, e.MyError)
}

func main() {
	fmt.Println("Ex.11.01")
	p1 := person{
		Last:  "Todd",
		First: "Todd",
	}
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf(string(bs))
	fmt.Println("Ex.11.02")
	fmt.Println("Ex.11.03")
	_, e := sqrt3(-10)
	if e != nil {
		log.Println(e)
	}
	fmt.Println("Ex.11.04")
	fmt.Println("Ex.11.05")
}

func toJSON(p interface{}) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		//return []byte{}, fmt.Errorf("new error found %v", err)
		return []byte{}, errors.New(fmt.Sprintf("new error found %v", err))
	}
	return bs, nil
}

func sqrt3(n float64) (float64, error) {
	if n < 0 {
		nma := fmt.Errorf("ex 03 error: %v", n)
		return 0, customError{1, nma}
	}
	return 42, nil
}
