/*Fix the race condition you created in exercise #4 by using package
atomic
code: https://github.com/GoesToEleven/go-programming
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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
	//	var atom int64

	var inc int64 = 0
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			//	runtime.Gosched()
			fmt.Println("counter : ", atomic.LoadInt64(&inc))
			wg.Done()
		}()
		//fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("GoRoutine numbers :", runtime.NumGoroutine())
	fmt.Println("Counter :", inc)

}
