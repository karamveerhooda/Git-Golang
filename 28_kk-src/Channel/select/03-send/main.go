// funny thing with sed channel. Here we send the data to channel after printing the channel value
// This behavior may be due to go func() will take separate goroutine than main goroutine
// and main go routine that have normal foo() call will work first than go fun routine
package main

import "fmt"

func main() {

	fmt.Println("hello")
	c := make(chan int)
	f := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		for i := 0; i < 5; i++ {
			fmt.Println(<-f)
		}
	}()

	foo(c)
	bar(f)
}

func foo(c chan int) {
	x := 1
	for i := 0; i < 5; i++ {
		select {
		case c <- x:
			x += x
			//		 fmt.Println("for channel C:", v)
		}
	}
}
func bar(f chan int) {
	x := 1
	for i := 0; i < 5; i++ {
		select {
		case f <- x:
			x += x
			//		 fmt.Println("for channel C:", v)
		}
	}
}
