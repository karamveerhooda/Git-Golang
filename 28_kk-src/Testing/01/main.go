/*Testing in Go, Three Takeaways
1.	Name of testing file should be in same package as of main program
2.	The test should be in file with name _test.go
3.	Tests must be in a func with name Testxxx(t *testing.t)
*/

package main

import "fmt"

func main() {
	fmt.Println("Sum of 2 + 3 is :", Mysum(2, 3))
	fmt.Println("Sum of 2 + 3 is :", Mysum(5, 5))
}

//Mysum func
func Mysum(sumvar ...int) int {
	sum := 0
	for _, v := range sumvar {
		sum += v
	}
	return sum + 1
}
