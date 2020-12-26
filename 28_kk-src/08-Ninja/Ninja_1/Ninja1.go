/*Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise -
remember to ask yourself what you need to do to EXPORT a value from a package.
solution: https://play.golang.org/p/8BK6PXj3aG */

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Phone int
}

func main() {

	user := []string{"karam", "hooda", "sapna", "choudhary", "surabhi", "sharma"}
	p1 := person{
		First: "karam",
		Last:  "hooda",
		Phone: 9711904500,
	}

	p2 := person{
		First: "surabhi",
		Last:  "sharma",
		Phone: 9711904343,
	}
	p3 := person{
		First: "sapna",
		Last:  "chawla",
		Phone: 9711904501,
	}

	people := []person{p1, p2, p3}
	fmt.Println(user)
	fmt.Println(people)

	//change user string in to JSON format using marshal
	//func Marshal(v interface{}) ([]byte, error)

	bs, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	bse, err1 := json.Marshal(people)
	if err1 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bse))
}
