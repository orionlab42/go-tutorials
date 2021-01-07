package main

import "fmt"

func main() {
	// arrays / not really used
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	// slices
	// y := type{value} // composite literal - values of the same type
	y := []int{4, 5, 7, 8, 42, 0}
	//fmt.Println(y)
	//fmt.Println(y[4])
	//fmt.Println(len(y))
	//fmt.Println(cap(y))
	for i, v := range y {
		fmt.Println(i, v)
	}
	//for i:=0; i < len(y); i++ {
	//	fmt.Println(i, y[i])
	//}
	y = append(y, 77, 88, 99)
	//fmt.Println(y)
	z := []int{234, 256, 298}
	y = append(y, z...)
	// having the ... after a slice it takes all its values
	fmt.Println(y)
	// deleting from a slice // basically appending the to the first part of the slice the values from the second part
	y = append(y[:6], y[8:]...)
	fmt.Println(y)
	newSlice := make([]int, 10, 100)
	fmt.Println(newSlice)
	//fmt.Println(len(new Slice))
	//fmt.Println(cap(newSlice))
	// multidimensional slice
	jb := []string{"James", "Bond", "Vanilla"}
	mp := []string{"Miss", "Penny", "Martini"}
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	// maps - key value stores - unordered list
	m := map[string]int{
		"James": 32,
		"Penny": 27,
	}
	fmt.Println(m)
	//fmt.Println(m["James"])
	//fmt.Println(m["Barnabas"])
	// check if it exist this value
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("There is no Barnabas.")
	}
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("There is no James.")
	}

	m["Todd"] = 34
	for k, v := range m {
		fmt.Println(k, v)
	}
	// this deletes it only if it exists
	if _, ok := m["James"]; ok {
		delete(m, "James")
	}
	fmt.Println(m)
}
