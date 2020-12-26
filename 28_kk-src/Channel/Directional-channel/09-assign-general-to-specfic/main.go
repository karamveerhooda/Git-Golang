package main

import "fmt"

func main() {

	xr := make(<-chan int) // it's bidrectional channel
	//Channel Send func
	x := make(chan int)
	xs := make(chan<- int)
	fmt.Printf("%T\t\n", x)
	fmt.Printf("%T\t\n", xr)
	fmt.Printf("%T\t\n", xs)

	xr = x
	//x = xr
}

//error comes :.\main.go:15:5: cannot use xs (type chan<- int) as type <-chan int in assignment
