// Package acdc asks if you are ready..
package acdc

// Sum adds an unlimited number of values of type int
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
