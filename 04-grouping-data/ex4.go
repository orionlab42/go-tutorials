package main

import "fmt"

func main() {
	fmt.Println("Ex4.1")
	//x := [5]int{1, 2, 3, 4, 5}
	////fmt.Println(x)
	//for i, v := range x {
	//	fmt.Println(i, v)
	//}
	//fmt.Printf("%T\t", x)
	fmt.Println("Ex4.2")
	//x := []int{1, 2, 3, 3, 3, 5, 8, 9, 10, 1}
	//fmt.Println(x)
	//fmt.Println(len(x))
	//for i, v := range x {
	//	fmt.Println(i, v)
	//}
	//fmt.Printf("%T\n", x)
	fmt.Println("Ex4.3")
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//fmt.Println(x[:5])
	//fmt.Println(x[5:])
	//fmt.Println(x[2:7])
	//fmt.Println(x[1:6])
	fmt.Println("Ex4.4")
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//x = append(x, 52)
	//fmt.Println(x)
	//x = append(x, 53, 54, 55)
	//fmt.Println(x)
	//y := []int{56, 57, 58, 59, 60}
	//x = append(x, y...)
	//fmt.Println(x)
	fmt.Println("Ex4.5")
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//x = append(x[:3], x[6:]...)
	//fmt.Println(x)
	fmt.Println("Ex4.6")
	//x := make([]int, 50, 50)
	//for i := 0; i < 50; i++ {
	//	x[i] = i + 1
	//	fmt.Println(i, x[i] )
	//}
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(cap(x))
	fmt.Println("Ex4.7")
	//jb := []string{"James", "Bond", "Shaken"}
	//mm := []string{"Miss", "Penny", "Hello"}
	//ss := [][]string{jb, mm}
	//for i, v := range ss {
	//	fmt.Println(i, v)
	//	for j, w := range v {
	//		fmt.Println(j, w)
	//	}
	//}
	fmt.Println("Ex4.8")
	fmt.Println("Ex4.9")
	fmt.Println("Ex4.10")
	m := map[string][]string{
		"bond_james": []string{"Shaken...", "Martinis", "Women"},
		"penny_miss": []string{"James Bond", "Literature", "Computers"},
		"dr_no":      []string{"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("This record is for %v\n", k)
		for i, val := range v {
			fmt.Printf("Index: %v\t Value: %v\n", i, val)
		}
	}

	m["todd_mctoad"] = []string{"computers", "alcohol"}

	for k, v := range m {
		fmt.Printf("This record is for %v\n", k)
		for i, val := range v {
			fmt.Printf("Index: %v\t Value: %v\n", i, val)
		}
	}

	if _, ok := m["todd_mctoad"]; ok {
		delete(m, "todd_mctoad")
	}

	for k, v := range m {
		fmt.Printf("This record is for %v\n", k)
		for i, val := range v {
			fmt.Printf("Index: %v\t Value: %v\n", i, val)
		}
	}
}
