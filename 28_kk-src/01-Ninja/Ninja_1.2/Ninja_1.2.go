//include ninja 1.2, 1.3,
package main

import "fmt"

var d = 42
var e = "James Bond"
var f = true

func main() {

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	s := fmt.Sprintf("\nd is %v\te is %v\tf is %v", d, e, f)
	fmt.Println(s)

}
