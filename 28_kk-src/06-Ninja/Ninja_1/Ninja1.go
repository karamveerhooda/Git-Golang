/*create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

package main

import "fmt"

func main() {
	x := 55
	foo(x)

	fmt.Println(x)
	fmt.Println(foo(x))

	str := "karam"

	fmt.Println(bar(str, x))
}

func foo(x int) int {
	x = 100
	return x
}
func bar(barstr string, z int) (string, int) {

	barstr = "hello"
	z = 11

	//fmt.Println(barstr)
	return barstr, z
}
