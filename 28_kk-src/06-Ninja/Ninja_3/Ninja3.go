/*Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
code: https://play.golang.org/p/XluEuUD0Nw*/

package main
import (
	"fmt"
)
func main() {

	defer deferfunc()

	fmt.Println("inside main")

	func(y int){
		fmt.Println("inside anyonymous func",y)
	}(22)
}

func deferfunc()  {
	fmt.Println("inside the deferfunc")
}

