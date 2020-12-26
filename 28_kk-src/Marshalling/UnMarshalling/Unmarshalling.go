//bcrypt is the process of encrypting the password
//
package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	// unmarshalling of above people value of person type
	//func Unmarshal(data []byte, v interface{}) error
	// unmarshaling is converting JSON format to type of format you have selected,
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`

	bs := []byte(s) // this is converting string of bytes in to slice of UNIT8 i.e. []UNIT8
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all data", people)

	//Bcrypt - Generating the hash password.
	passw := `Password123`
	//var s []byte
	bs, err1 := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err1 != nil {
		fmt.Printf("password not generated")
	}

	fmt.Println("Encrypt password is :", string(bs))
	passw1 := `Password13`
	err1 = bcrypt.CompareHashAndPassword(bs, []byte(passw1))
	if err1 != nil {
		fmt.Printf("Not logged in")
		return
	}
	fmt.Println("You are logged in")

}
