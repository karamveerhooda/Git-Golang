/*create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type []int into your func (unfurl the []int)
	returns the sum of all values of type int passed in
create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
code: https://play.golang.org/p/B0yRxtBQPD */
package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	fmt.Println("Sum of slice from foo :", foo(x...))

	y := []int{2, 4, 6}

	fmt.Println("Sum of slice from bar :", bar(y))

}

func bar(z []int) int {
	var sum int

	for _, v := range z {
		sum += v
	}

	return sum
}

func foo(y ...int) int {
	var sum int

	for _, v := range y {
		sum += v
	}

	return sum
}
