package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//n, err := fmt.Println("Checking errors: ")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(n)
	//f, e := os.Create("names.txt")
	//if e != nil {
	//	fmt.Println(e)
	//	return
	//}
	//defer f.Close()
	//r := strings.NewReader("Whatsapp")
	//io.Copy(f, r)
	f, e := os.Open("names.txt")
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	bs, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(bs))
	_, e = os.Open("no-file.txt")
	if e != nil {
		//fmt.Println("Error happened: ",e)
		//log.Println("Error happened: ",e)
		//log.Fatalln(e)  // no defers will be ran it will right away exit
		//log.Panicln(e) // defers will be ran it will gradually exit
		panic(e) // here there is no logging just printing
	}
}
