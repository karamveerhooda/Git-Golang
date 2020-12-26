// 1. To find the even numbers from a given slice by using
package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//s := sum(ii...)
	//	fmt.Println(s)

	s2 := even(sum, ii...) // s2 := foo(bar)
	fmt.Println(s2)        //print s2
}
func sum(xii ...int) int { // bar()
	//	var sii []int
	s := 0
	for _, v := range xii {
		if v%2 == 0 {
			s += v
		}
	}
	fmt.Println("sum of even numbers :", s)
	return s
}

// in even func we pass function as a first parameter with xii as a variadic int parameter (xii ...int) with return value of int and second parameter vi as a variadic paramter of int
func even(f func(xii ...int) int, vi ...int) int { // func foo(f func())

	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	fmt.Println(yi)
	return f(yi...)
}
