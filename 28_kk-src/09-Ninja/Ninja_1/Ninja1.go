/*Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists
code: https://github.com/GoesToEleven/go-programming
video: 148
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutine number :", runtime.NumGoroutine())
	fmt.Println("CPU number :", runtime.NumCPU())
	wg.Add(2)
	go outing()
	go whereouting()
	fmt.Println("In between Goroutine number :", runtime.NumGoroutine())
	fmt.Println("In between CPU number :", runtime.NumCPU())
	wg.Wait()
	fmt.Println("End Goroutine number :", runtime.NumGoroutine())
	fmt.Println("end CPU number :", runtime.NumCPU())
}

func outing() {
	for i := 0; i <= 5; i++ {
		fmt.Println("we want to go out")
	}
	wg.Done()
}
func whereouting() {
	for i := 0; i <= 5; i++ {
		fmt.Println("we don't know where to go out")
	}
	wg.Done()
}
