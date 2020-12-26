package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	/*	f, err := os.Create("myfile.txt")
		if err != nil {
			fmt.Println(err)
		}

		defer f.Close()
		log.SetOutput(f)*/
	// SO if deferred func is there it will not execute after commiting
	//fatal error in program
	defer foo()

	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	//defer f2.Close()
	//log.SetOutput(f)
}
func foo() {
	fmt.Println("IN foo and  it should not execute")
}
