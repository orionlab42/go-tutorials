package main

import (
	"fmt"
	"sort"
)

type person3 struct {
	first string
	age   int
}

//func (p person3) String() string {
//	return fmt.Sprintf("From string method %s: %d", p.first, p.age)
//}

type ByAge []person3

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []person3

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	fmt.Println("Sort")
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr.No"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("Sort custom")
	p1 := person3{"James", 32}
	p2 := person3{"Moneypenny", 27}
	p3 := person3{"Q", 64}
	p4 := person3{"A", 56}

	people := []person3{p1, p2, p3, p4}
	fmt.Println(people)
	//s1 := p1.String()
	//fmt.Println(s1)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
