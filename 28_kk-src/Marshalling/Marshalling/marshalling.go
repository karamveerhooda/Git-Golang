package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "karam",
		Last:  "hooda",
		Age:   39,
	}
	p2 := person{
		First: "punam",
		Last:  "khokhar",
		Age:   37,
	}
	fmt.Println(p1, p2)

	//func Marshal(v interface{}) ([]byte, error)

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
