/*Hands-on exercise #5
	create a type SQUARE  --means create a struct
	create a type CIRCLE
	attach a method to each that calculates AREA and returns it
		circle area= π r 2
		square area = L * W
	create a type SHAPE that defines an interface as anything that has the AREA method
	create a func INFO which takes type shape and then prints the area
	create a value of type square
	create a value of type circle
	use func info to print the area of square
	use func info to print the area of circle
code: https://play.golang.org/p/NGGikHNvHv

*/
package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	c1 := circle{
		radius: 5,
	}
	s1 := square{
		side: 5,
	}

	info(c1)
	info(s1)
}
