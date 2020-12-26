// This program uses three channels , two channels used for send and recieve
//and third channel used to send data from first two channel.
//USes simple waitgroup, and for range to extract the value from channel

package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fan := make(chan int)

	go send(even, odd)
	go recieve(even, odd, fan)
	for v := range fan {
		fmt.Println(v)
	}
}

func send(even, odd chan<- int) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func recieve(even, odd <-chan int, fan chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fan <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fan <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fan)
}
