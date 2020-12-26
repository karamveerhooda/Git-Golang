/*Fan-In ->  taking value from two channel and merge it in to one channel
step 1: Call a func which internally call two same func with
		string paramter.
step 2: The return value from step 1 func is receive channel.
step 3: The internal two func(basically single func as both are same)
		is having a channel receiving a data from normal for loop including the
		string which is passed in to it.
step 4: In step 1 func that have receive only return channel, will recieve the
		value in another channel using two GO routine because this func have two func parameters
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := fanIn(boring("Tv"), boring("Doordarshan"))

	//for v := range c {
	//		fmt.Println(v)
	//	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
func boring(tv string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", tv, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c

}
func fanIn(input1, Input2 <-chan string) <-chan string {

	c := make(chan string)
	go func() {

		for {
			c <- <-input1
		}
	}()
	go func() {

		for {
			c <- <-Input2
		}
	}()
	return c
}
