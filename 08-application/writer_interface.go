package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Writer interface")
	fmt.Fprintln(os.Stdout, "Writer interface2")
	io.WriteString(os.Stdout, "Writer interface3")
}
