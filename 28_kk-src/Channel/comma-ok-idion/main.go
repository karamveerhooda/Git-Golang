/*This is comma ok idiom, to check what happen when we use v,ok idiom on channel*/
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 10
		close(c)
	}()
	v, ok := <-c

	fmt.Println(v, ok)
	v, ok = <-c

	fmt.Println(v, ok)

}
