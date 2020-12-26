package main

import (
	"errors"
	"fmt"
	"log"
)

var errormessage = errors.New("Entered number for sqrt is wrong")

func main() {

	fmt.Printf("\n%T\n :", errormessage)
	//fmt.Printf("\n%T :", fmt

	//	_, err := fmt.Println()
	_, err := sqrt(-10)
	fmt.Printf("\n%T\n :", err)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errormessage
	}
	return 42, nil
}
