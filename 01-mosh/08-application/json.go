package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   52,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   32,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	// Marshaling
	bs1, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs1))
	// Unmarshaling
	s := `[{"First":"James","Last":"Bond","Age":52},{"First":"Miss","Last":"Moneypenny","Age":32}]`
	bs2 := []byte(s)
	var peopleUnmarshal []person2
	fmt.Println(bs2)
	fmt.Printf("%T\n", bs2)
	err = json.Unmarshal(bs2, &peopleUnmarshal)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(peopleUnmarshal)
	for i, v := range peopleUnmarshal {
		fmt.Println("\n Person number: ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
