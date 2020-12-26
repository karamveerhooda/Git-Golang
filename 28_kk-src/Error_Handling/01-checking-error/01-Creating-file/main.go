/**/
package main

import (
	"fmt"
	"os"
)

func main() {
	//file creation

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var n int
	n, err = f.WriteString("karamhooda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	//strings.NewReader
}
