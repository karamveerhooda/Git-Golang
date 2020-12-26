package main

import "testing"

type test struct {
	data []int
	res  int
}

//Testmysum function of main
func TestMysum(t *testing.T) {

	tests := []test{
		test{[]int{1, 2}, 3},
		test{[]int{2, 2}, 4},
		test{[]int{3, 2}, 5},
		test{[]int{4, 2}, 6},
	}

	for _, v := range tests {
		x := Mysum(v.data...)
		if x != v.res {
			t.Error("Expected", v.res, "Got", x)
		}
	}
}
