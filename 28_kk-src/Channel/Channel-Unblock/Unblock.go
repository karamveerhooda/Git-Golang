/*Channel are used to make th ecode concurrent and works in a way
of sending and recieving at the same time. So if receiver is not rady then
channel will block the send item and gives error ALL THREADS ARE ASLEEP*/
package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan int)
	go func() { // This will allow to print the chanvar var
		//because this goroutine and main goroutine go hand in hand
		//and func goroutine pass the baton hand in hand to main func goroutine.
		//So there is no delay.
		chanvar <- 10
	}()

	fmt.Println(<-chanvar)
	chanvar2 := make(chan int, 1)
	chanvar2 <- 20

	fmt.Println(<-chanvar2) // This will print the value of chanvar2
	//because channel have 1 buffer (chan int,1) and can wait for passing the baton only
	//to one receiver
}
