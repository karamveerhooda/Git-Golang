package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	i := 0
	fmt.Println("value of i before calling c: ", i)
	y := c()
	fmt.Println("value of i after calling c: ", y)

	_, err := os.Open("no-file.txt")

	if err != nil {
		panic(err)
	}
}

//1. Deferred function's argument are evaluated when the defer
// statement is evaluated
//2.Deferred function calls are executed in LIFO manner
func foo() {

	i := 0
	defer fmt.Println("value of i is :", i)
	i++
	fmt.Println("value of I after one increment :", i)
	return

}

//3. Deferred function may read and assign to the returning function's
// named return values
func c() (i int) {

	defer func() {
		i++
	}()
	return 1
}
