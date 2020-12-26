package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file")
	if err != nil {
		log.Println("error happened ", err)
		fmt.Println(err)
	}
	defer f2.Close()
	//log.SetOutput(f)
}
