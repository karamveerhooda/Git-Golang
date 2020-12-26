package main

import "fmt"

type chess int

var x chess

func main() {

	fmt.Println(x)
	fmt.Printf("%T\t\n", x)

	x = 42
	fmt.Println(x)

}
