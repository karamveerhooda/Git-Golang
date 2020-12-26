package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan<- int) // it's bidrectional channel
	//Channel Send func

	chanvar <- 10
	go func() {
		fmt.Println(<-chanvar)
	}()
}
