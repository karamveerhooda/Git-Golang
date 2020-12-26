/*Hands-on exercise #3
	Using goroutines, create an incrementer program
		have a variable to hold the incrementer value
		launch a bunch of goroutines
			each goroutine should
				read the incrementer value
					store it in a new variable
				yield the processor with runtime.Gosched()
				increment the new variable
				write the value in the new variable back to the incrementer variable
	use waitgroups to wait for all of your goroutines to finish
	the above will create a race condition.
	Prove that it is a race condition by using the -race flag
	if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
code: https://github.com/GoesToEleven/go-programming
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// generating race condition, declare one counter variable inc with value 0 and then
	//declare one goroutne and assign the value of inc in to new variable incval
	//increment incaval
	//and assign back incval to inc
	//this will generate the race condition.
	//check the race condition with command go run -race ninja3.go command
	//Now check the subroutine also.

	var wg sync.WaitGroup
	inc := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			incval := inc
			runtime.Gosched()
			incval++
			inc = incval
			wg.Done()
			fmt.Println("counter :", inc)
			//	fmt.Println("GoRoutine numbers inside func() :", runtime.NumGoroutine())
		}()
		//fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())
	fmt.Println("Counter :", inc)

}
