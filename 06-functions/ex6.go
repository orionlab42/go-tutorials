package main

import "fmt"

func main() {
	fmt.Println("Ex.06.01")
	fooo1 := fooo()
	fmt.Println("Returned by fooo: ", fooo1)
	baaar1, baaar2 := baaar()
	fmt.Println("Returned by baaar: ", baaar1, baaar2)
	fmt.Println("Ex.06.02")

	fmt.Println("Ex.06.03")
}

//fmt.Println("Ex.06.01")
func fooo() int {
	return 200
}
func baaar() (int, string) {
	return 400, "Yesss"
}

//fmt.Println("Ex.06.02")
