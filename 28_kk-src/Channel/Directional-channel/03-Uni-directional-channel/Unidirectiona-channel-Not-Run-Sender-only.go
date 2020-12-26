package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan<- int, 2) // send only channel
	chanvar <- 10
	chanvar <- 20

	fmt.Println(<-chanvar) // Thi will give error  <-chanvar (receive from send-only type chan<- int)
	fmt.Println(<-chanvar)

	fmt.Printf("%T\t", chanvar)
}
