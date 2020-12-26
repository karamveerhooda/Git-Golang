//all exercises
package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		fmt.Printf("%v\t", x)
	}

	z := 1
	for x := byte('A'); x <= byte('Z'); x++ {
		fmt.Println(z)
		fmt.Println()
		for y := 0; y <= 2; y++ {
			fmt.Printf("\t%U %q", x, x)
			fmt.Println()
		}
		z++
	}

	bd := 1985
	for bd <= 2020 {
		fmt.Printf("%d\t", bd)
		bd++
	}
	fmt.Println()
	bd = 1985
	for {
		fmt.Printf("%d\t", bd)
		bd++
		if bd > 2020 {
			break
		}

	}

	for i := 10; i <= 100; i++ {
		var x int = i % 4
		fmt.Printf("when %v is divided by 4 then remainder is %v\n", i, x)

		if x >= 3 {
			fmt.Println("I am ok")
		} else {
			fmt.Println("I am less than 3")
		}
	}

	//	var x bool
	switch {
	case false:
		fmt.Println("shouldn't print")
	case true:
		fmt.Println("should print")
	default:
		fmt.Println("I am true")
	}

	var favsport string = "cricket"

	switch favsport {
	case "cricket":
		fmt.Println("Cricket is my favorite gmae")
	case "football":
		fmt.Println("Football is my favorite one")
	default:
		fmt.Println("gulli danda")
	}
	//how many 6 digit number have atleast 5 9's

}
