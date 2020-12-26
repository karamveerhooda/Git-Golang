/*Hands-on exercise #5
check channel is closed or closed
Show the comma ok idiom starting with this code.
solution: https://play.golang.org/p/qh2ywLB5OG
video: 168
*/package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 44
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}
