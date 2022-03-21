package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate error occured: %v, %v, %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt2(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt2(n float64) (float64, error) {
	if n < 0 {
		nma := fmt.Errorf("error as a nma: %v", n)
		return 0, norgateMathError{"latitude", "altitude", nma}
	}
	return 42, nil
}
