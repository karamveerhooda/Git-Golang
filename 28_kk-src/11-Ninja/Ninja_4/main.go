/*Hands-on exercise #4
Starting with this code, use the sqrt.Error struct as a value of type error.
If you would like, use these numbers for your
lat "50.2289 N"
long "99.4656 W"
solution:
https://play.golang.org/p/nsRxbDfkCh
video: 179*/
package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//lat "50.2289 N" long "99.4656 W"

		//e := errors.New("more cofee needed")
		//return math.Sqrt(f), sqrtError{"50.2289 N", "99.4656 W", e}
		return math.Sqrt(f), fmt.Errorf(fmt.Sprintf("errorcoming %v", f))
	}
	return 42, nil
}
