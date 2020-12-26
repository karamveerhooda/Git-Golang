/*Channel are used to make th ecode concurrent and works in a way
of sending and recieving at the same time. So if receiver is not rady then
channel will block the send item and gives error ALL THREADS ARE ASLEEP*/
package main

import (
	"fmt"
)

func main() {

	chanvar := make(chan int, 2)
	chanvar <- 10
	chanvar <- 20

	fmt.Println(<-chanvar)
	fmt.Println(<-chanvar)
}
