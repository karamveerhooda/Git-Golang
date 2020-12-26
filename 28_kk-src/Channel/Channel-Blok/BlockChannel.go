/*Channel are used to make th ecode concurrent and works in a way
of sending and recieving at the same time. So if receiver is not rady then
channel will block the send item and gives error ALL THREADS ARE ASLEEP*/
package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan int)

	chanvar <- 10

	fmt.Println(<-chanvar) // This will not print the value of chanvar
	//because value is not received hand in hand from sender.
	//Causing error All goroutines are asleep - deadlock
}
