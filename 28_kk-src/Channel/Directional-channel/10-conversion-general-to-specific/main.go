package main

import "fmt"

func main() {

	x := make(chan int)
	xs := make(chan<- int)
	xr := make(<-chan int)

	fmt.Printf("%T\t\n", x)
	fmt.Printf("%T\t\n", xr)
	fmt.Printf("%T\t\n", xs)

	fmt.Println("------------")
	fmt.Printf("%T\t\n", (chan<- int)(x))
	fmt.Printf("%T\t\n", (<-chan int)(x))
	//	fmt.Printf("%T\t\n", (chan int)(xs))

	//xr = x
	//x = xr
}

//error comes :.\main.go:15:5: cannot use xs (type chan<- int) as type <-chan int in assignment
