package main

import "fmt"

func main() {
	//for i := 1; i <= 10000; i++ {
	//	if i%10 == 0 {
	//		fmt.Printf("%v\n", i)
	//	} else {
	//		fmt.Printf("%v\t", i)
	//	}
	//}
	//for i:=48; i<=90; i++ {
	//	fmt.Printf("%v\n", i)
	//	for j:=0; j<3;j++ {
	//		fmt.Printf("\t\t%#U and %x\n", i, i)
	//	}
	//}
	//i:=1989
	//for i<=2020 {
	//	fmt.Printf("I am alive yeahhhh and the year is: %v\n", i)
	//	i++
	//}

	//i:=1989
	//for {
	//	fmt.Printf("I am alive yeahhhh and the year is: %v\n", i)
	//	if i >= 2020 {
	//		break
	//	}
	//	i++
	//}

	//for i:= 10; i <=100; i++ {
	//	l := i%4
	//	fmt.Printf("For %v the modulus is %v.\n", i, l)
	//	if i%4 == 0 {
	//		fmt.Println(i, " is dividable with 4.")
	//	} else if i%4 == 1 {
	//		fmt.Println("Nope")
	//	} else if i%4 == 2 {
	//		fmt.Println("Nope nope")
	//	} else {
	//		fmt.Println("Nope nope nope")
	//	}
	//}

	for i := 10; i <= 100; i++ {
		switch {
		case i%4 == 0:
			fmt.Println(i, " is dividable with 4.")
		case i%4 == 1:
			fmt.Println("Nope")
		case i%4 == 2:
			fmt.Println("Nope nope")
		default:
			fmt.Println("Nope nope nope")
		}
	}
	//
	//favSport := "read"
	//switch favSport {
	//case "surf":
	//	fmt.Println("Yeahh")
	//case "run":
	//	fmt.Println("oh noo..anyway")
	//default:
	//	fmt.Println("you dont know")
	//}
}
