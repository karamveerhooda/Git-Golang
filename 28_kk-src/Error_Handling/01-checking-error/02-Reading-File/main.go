package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Open file error :", err)
	}

	defer f.Close()
	s, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read file error :", err)
	}
	fmt.Println(string(s))
}
