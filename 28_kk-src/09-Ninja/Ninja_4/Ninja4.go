/*Hands-on exercise #4
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
code: https://github.com/GoesToEleven/go-programming
video: 151
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
	var mu sync.Mutex
	inc := 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			incval := inc
			//runtime.Gosched()
			incval++
			inc = incval
			fmt.Println("counter :", inc)
			mu.Unlock()
			wg.Done()
			//	fmt.Println("GoRoutine numbers inside func() :", runtime.NumGoroutine())

		}()
		//fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())
	fmt.Println("Counter :", inc)

}
