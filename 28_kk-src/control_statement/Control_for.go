package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	x := 1
	for {
		if x > 5 {
			break
		}
		fmt.Println(x)
		x++
	}
	//use continue - find odd numbers from 1 to 100
	a := 1
	for {
		a++
		if a > 10 {
			break
		}
		if a%2 == 0 {
			continue
		}
		fmt.Printf("%d\t", a)
		fmt.Println()
		fmt.Println()

	}
	for k := 33; k <= 123; k++ {
		fmt.Printf("%v\t%#x\t%#U\n", k, k, k)
	}

}
