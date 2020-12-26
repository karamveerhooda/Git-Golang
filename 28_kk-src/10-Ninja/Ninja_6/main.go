/*Hands-on exercise #6
write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them
solution: https://play.golang.org/p/096Lk1BR7o
video: 169
*/

package main

import "fmt"

func main() {
	c1 := make(chan int)

	go send(c1)
	recieve(c1)
	fmt.Println("about to exit")
}
func send(c1 chan int) {
	for i := 0; i < 1000; i++ {
		c1 <- i
	}

	close(c1)
}

func recieve(c1 chan int) {
	for v := range c1 {
		fmt.Println(v)
	}
}
