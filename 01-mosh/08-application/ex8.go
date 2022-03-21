package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//type user struct {
//	First string
//	Age int
//}

type sayings struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge2 []user

func (a ByAge2) Len() int           { return len(a) }
func (a ByAge2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge2) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (b ByLast) Len() int           { return len(b) }
func (b ByLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }

func main() {
	fmt.Println("Ex.08.01")
	//u1 := user{
	//	First: "James",
	//	Age: 62,
	//}
	//u2 := user{
	//	First: "Todd",
	//	Age: 27,
	//}
	//users := []user{u1, u2}
	//bs, err := json.Marshal(users)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(bs)
	//fmt.Println(string(bs))

	fmt.Println("Ex.08.02")
	jsonToConvert := `[{"First":"James", "Last":"Bond","Age":32,"Sayings":["Shaken, but not stirred","Youth is no ..","Bla bla bla"]}, {"First":"Todd", "Last":"McLoad","Age":55,"Sayings":["Ease of programming","Obsessed with Bond","Bla bla bla"]}]`
	var sayingsCollection []sayings
	err := json.Unmarshal([]byte(jsonToConvert), &sayingsCollection)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range sayingsCollection {
		fmt.Println("Person nr.", i)
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Printf("\t%v\n", val)
		}
	}
	fmt.Println("Ex.08.03")
	user1 := user{
		First: "James",
		Last:  "Bond",
		Age:   35,
		Sayings: []string{
			"Shaken, but not stirred",
			"Youth is no ..",
			"Ala bla bla",
			"In his majesty's royal service"},
	}
	user2 := user{
		First: "Todd",
		Last:  "McLoad",
		Age:   20,
		Sayings: []string{
			"Ease of programming",
			"Obsessed with Bond",
			"Ala bla bla"},
	}
	users := []user{user1, user2}
	//fmt.Println(users)

	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ex.08.04")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 55}
	xs := []string{"random", "rainbow", "in", "torpedo", "summers", "under", "gallantry"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println("Ex.08.05")
	fmt.Println(users)
	sort.Sort(ByAge2(users))
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)
	sort.Strings(user1.Sayings)
	fmt.Println(users)
}
