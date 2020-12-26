/*Hands-on exercise #7
write a program that
launches 10 goroutines
each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
solutions:
https://play.golang.org/p/R-zqsKS03P
https://play.golang.org/p/quWnlwzs2z
student solutions:
https://twitter.com/mannion
https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	c1 := make(chan int)

	for i := 0; i <= 10; i++ {
		go addgoroutine(c1)
		fmt.Println("num routine :", runtime.NumGoroutine())

	}

	//close(c1)

	recieve(c1)
	fmt.Println("About to exit ")
}

func addgoroutine(c1 chan<- int) {
	for i := 0; i <= 100; i++ {
		c1 <- i
	}
	//close(c1)
}
func recieve(c1 <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c1)
	}
}
