package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{0, 0}, 0},
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 45},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Error("Expected", sum, "Got", v.answer)
		}
	}
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
