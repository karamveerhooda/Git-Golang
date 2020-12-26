//Package Dog is being imported in Ninja_12 Package main.go program
package Dog

import (
	"fmt"
)

//Years to be calculated
func Years(hy *int) {
	//var dogyears int
	dogyears := *hy * 7
	fmt.Println("Dog years for given Human years are :", dogyears)
}

func Foo() {
	fmt.Println("in foo")
}
