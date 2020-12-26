package main

import (
	"fmt"
)

//declaring variable outside the func bdy you need var
//variable with the identifier z is of type int
//Each element of such a variable or value is set to the zero value for its
// type: false for booleans, 0 for numeric types, "" for strings,
//and nil for pointers, functions, interfaces, slices, channels, and maps.

// string can be declared between "" or ''(backquote)
var z = 911

type hotdog int

var a hotdog

func main() {
	n, _ := fmt.Println("Hello world and Avoid chinese")
	fmt.Println(n)

	x := 50
	fmt.Println(x)

	y := 50 + 34
	fmt.Println(y)

	fmt.Printf("\n%T\n", z)
	fmt.Printf("%b\n", z)
	fmt.Printf("%#x\a\n", z)

	fmt.Printf("%T\t %b\t %x\t\n", z, z, z)
	s := fmt.Sprintf("%v\t%#x\t %b\t %x\t", z, z, z, z)
	fmt.Println(s)

	a = 10

	fmt.Printf("\n%T\n", a)
	//	z = a  not allowed
	//Below this is allowed
	z = int(a) // this is allowed , this is not called casting in GOLANG
	// this is called conversions
	fmt.Println(z)

}

//every value is also type of interface{}
//fmt.println()
//above line read as package.Identifier.
