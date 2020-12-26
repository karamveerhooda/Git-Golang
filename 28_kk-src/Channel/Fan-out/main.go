package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanoutIN(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

}
func populate(c1 chan int) {

	for i := 0; i < 10; i++ {
		c1 <- i
	}
	close(c1)
}
func fanoutIN(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutine = 10
	wg.Add(goroutine)

	for i := 0; i < goroutine; i++ {
		go func() {
			for v := range c1 {
				//func(v2 int) {
				c2 <- timeconsumingWork(v)
				//}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

// This function helps in getting the random number till 50
func timeconsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(50)
}
