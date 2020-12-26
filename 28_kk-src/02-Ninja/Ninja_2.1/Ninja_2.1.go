//Ninja 2.1, 2.2, 2.3, 2.4
package main

import "fmt"

const (
	untyp     = 12.1
	typ   int = 12
)

var shiftvar int

const (
	year1 = iota + 2016
	year2 = iota + 2016
	year3 = iota + 2016
	year4 = iota + 2016
)

func main() {

	var exer4 int = 100
	var num int = 100
	fmt.Printf("num in binary %b\n", num)
	fmt.Printf("num in hexa %#x\n", num)
	fmt.Printf("num in decimal %d\n", num)

	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 42)
	d := (41 < 42)
	e := (41 != 41)

	fmt.Println(a, b, c, d, e)
	fmt.Printf("%T\t Untyped var\n", untyp)
	fmt.Printf("%T\t typed var\n", typ)

	fmt.Printf("%d in decimal\t%#x in hexa\t%b in binary\n", exer4, exer4, exer4)

	shiftvar = exer4 >> 1
	fmt.Printf("%d in decimal\t%#x in hexa\t%b in binary\n", shiftvar, shiftvar, shiftvar)
	// exer 5 ninja 2(2.5)
	str := `I have to fiish "Golang" ASAP and also finish "sysops" and cloud developer 
certification
parallely prepare for IELTS to get the 6
band atleast` //` ` this is calle he raw string literals

	fmt.Println(str)

	// exercise 2.6 Ninja.
	fmt.Println(year1, year2, year3, year4)

	//
}
