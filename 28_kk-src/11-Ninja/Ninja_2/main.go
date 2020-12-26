/*Hands-on exercise #2
Start with this code. Create a custom error message using “fmt.Errorf”.
solution:
https://play.golang.org/p/HugU4HJEEO
https://play.golang.org/p/NII-lmGasj
https://play.golang.org/p/Vo5kIoR-sG
video: 177*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

}
func toJSON(a interface{}) ([]byte, error) {
	bs, error := json.Marshal(a)
	if error == nil {
		//return []byte{}, fmt.Errorf("There was an error in toJSON : %v", error)
		return []byte{}, errors.New("There was an error in toJSON : %v")

	}
	return bs, nil
}
