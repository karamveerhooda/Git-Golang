package main

import "fmt"

func main() {

	xr := make(<-chan int) // it's bidrectional channel
	//Channel Send func
	x := make(chan int)
	xs := make(chan<- int)
	fmt.Printf("%t\t", x)
	fmt.Printf("%t\t", xr)
	fmt.Printf("%t\t", xs)

	xr = xs
	//x = xr

}

//error comes :.\main.go:15:5: cannot use xs (type chan<- int) as type <-chan int in assignment
