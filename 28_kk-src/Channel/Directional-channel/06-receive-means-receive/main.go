package main

import (
	"fmt"
)

func main() {

	chanvar := make(<-chan int) // it's bidrectional channel
	//Channel Send func

	chanvar <- 10
	go func() {
		fmt.Println(<-chanvar)
	}()
}

//error comes : .\main.go:12:10: invalid operation: chanvar <- 10 (send to receive-only type <-chan int)
