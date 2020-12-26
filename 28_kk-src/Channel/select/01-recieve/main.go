package main

import "fmt"

func main() {

	fmt.Println("hello")
	c := make(chan int)
	f := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			f <- i
		}
	}()

	foo(c, f)
}

func foo(c, f chan int) {
	for i := 0; i < 10; i++ {
		select {
		case v := <-c:
			fmt.Println("for channel C:", v)
		case v := <-f:
			fmt.Println("for channel f:", v)
		}
	}
}
