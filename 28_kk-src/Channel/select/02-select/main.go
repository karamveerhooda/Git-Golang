//select statement used to pulling off the values from the channel, If you have
//different channel and you want to take the value off then use Select
//in thsi program make even, odd and quit channels to pull the respecive channel using SELECT
package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send channel
	go send(even, odd, quit)

	receive(even, odd, quit)
}
func send(e, o chan<- int, q chan<- bool) {

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
		//		q <- 0
	}
	close(q)
}
func receive(e, o <-chan int, q <-chan bool) {

	for {
		select {
		case v := <-e:
			fmt.Println("Even number :", v)
		case v := <-o:
			fmt.Println("Odd number  :", v)
		case v, ok := <-q:
			if v == ok {
				fmt.Println("Quit :", v)
			} else {
				fmt.Println(v)
			}
			return
		}
	}
}
