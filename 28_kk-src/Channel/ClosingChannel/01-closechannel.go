/*Closing the channel*/
package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//	fmt.Println("hello")
	//Send value to channel
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//Recieve channel value
	func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	fmt.Println("exit")

}
