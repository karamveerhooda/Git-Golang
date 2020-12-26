//Example for GODOC documentation
package main

import (
	"fmt"

	"github.com/GoesToEleven/GolangTraining/28_kk-src/Ninja_12/01/Dog"
	//"C:/Users/karam/OneDrive/Desktop/AWS/GOLANG/goworkspace/src/github.com/GoesToEleven/GolangTraining/28_kk-src/Ninja_12/01/Dog/"
)

var humanyears int

func main() {
	fmt.Println("Enter the human years :")
	fmt.Scanf("%d", &humanyears)
	//Exported function
	Dog.Years(&humanyears)

	//	fmt.Println("hello")
}

//Years is Exported function
