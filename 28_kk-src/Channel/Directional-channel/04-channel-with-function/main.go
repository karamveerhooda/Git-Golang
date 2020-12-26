package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan int) // it's bidrectional channel
	//Channel Send func
	go foo(chanvar)

	//Channel Receiver
	bar(chanvar)
	fmt.Println("About to exit")

	//fmt.Printf("%T\t", chanvar)
}
func foo(chanvar chan<- int) {
	chanvar <- 10
}
func bar(chanvar <-chan int) {
	fmt.Println(<-chanvar)
}
