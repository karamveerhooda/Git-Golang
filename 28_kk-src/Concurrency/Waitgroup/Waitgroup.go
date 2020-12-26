package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Archs:\t\t", runtime.GOARCH)
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("---------------------------------")

	wg.Add(1)
	go foo() // if I comment this line , you will receive fatal error - fatal error: all goroutines are asleep - deadlock!
	woo()
	fmt.Println("CPU's:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("---------------------------------")
	wg.Wait()

}

func foo() {
	fmt.Println("In Foo")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func woo() {
	fmt.Println("in woo")
	for j := 10; j >= 0; j-- {
		fmt.Println(j)
	}
}
