/*
 Q: How should interfaces program work?
 A: --First declare the interface
	-- Define and declare each interface func separately as per the your own native requirement like below define area() method separately for rectangle and circle.
	   Make sure you are using the right type (of struct) before defining the interfaces method like func (c circle) area() float64{}.
	 -- Then declare a func that uses interfaces as a paramter and show what is done by interfaces in built function. e.g. g.area() or g.parameter()
	 -- Now in Main func passes the different variable declare for struct to test the different functionality of func in interfaces. for e.g in below main fun
		passing r1, reec1, r2, rec2 in testing method measure.
*/

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	parameter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) parameter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) parameter() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area ", g.area())
	fmt.Println("Parameter ", g.parameter())
}

func main() {

	r1 :=
		circle{
			radius: 5,
		}

	rec1 :=
		rect{
			width: 5,

			height: 5,
		}
	r2 :=
		circle{
			radius: 5,
		}

	rec2 :=
		rect{
			width: 5,

			height: 5,
		}
	measure(r1)
	measure(rec1)
	measure(r2)
	measure(rec2)

	//Anoynmos function

	func(x int) {
		fmt.Println("In Anonymous function", x)
	}(33)

}
