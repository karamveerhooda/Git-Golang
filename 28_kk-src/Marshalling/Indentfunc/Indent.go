package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type person struct {
	First string
	Last  string
	//	Age   int
}

func main() {

	p1 := []person{
		{"karam", "Hooda"},
		{"karam", "Hooda"},
		{"karam", "Hooda"},
		{"karam", "Hooda"},
		{"karam", "Hooda"},
	}

	fmt.Println(p1)

	//func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
	//func Marshal(v interface{}) ([]byte, error)
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	var out bytes.Buffer

	err = json.Indent(&out, b, "-", "\t")
	if err != nil {
		fmt.Println(err)
	}
	out.WriteTo(os.Stdout)

	xi := []int{2, 556, 22, 3, 4, 5, 887, 5, 55, 332}
	xs := []string{"f", "karam", "hooda", "punam", "khokhar", "tejus hoods", "ranishka Hooda", "a", "BGG", "awsed"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("----------------------------------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
