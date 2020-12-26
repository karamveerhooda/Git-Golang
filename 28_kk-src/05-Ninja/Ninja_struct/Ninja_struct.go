package main

import "fmt"

/*Exercise1 :Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

type myperson struct {
	firstname string
	lastname  string
	flavor    []string
}

func main() {

	myperson1 := myperson{
		firstname: `karam`,
		lastname:  `hooda`,
		flavor:    []string{`vanila`, `strawberry`},
	}
	myperson2 := myperson{
		firstname: `punam`,
		lastname:  `khokhar`,
		flavor:    []string{`vanila`, `strawberry`},
	}

	/*person := person{
		firstname: "punam",
		lastname:  "khokhar",
		flavor:    "strawberry",
	}*/
	fmt.Println(myperson1.firstname)
	fmt.Println(myperson1.lastname)
	for i, v := range myperson1.flavor {
		fmt.Println(i, v)
	}
	fmt.Println(myperson2.firstname)
	fmt.Println(myperson2.lastname)
	for i, v := range myperson2.flavor {
		fmt.Println(i, v)
	}
	fmt.Println()
	/* Q2:Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	   Access each value in the map. Print out the values, ranging over the slice.
	*/

	mymap := map[string]myperson{
		myperson1.lastname: myperson1,
		myperson2.lastname: myperson2,
	}

	for _, v := range mymap {
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for i, val := range v.flavor {
			fmt.Println(i, val)
		}
		fmt.Println("-----------")
	}

	/* Q4: Make the anoynmous struct with fields Map and slice and print them*/

	s := struct {
		place  string
		map1   map[string]int
		slice1 []string
	}{
		place: "Gurgaon",
		map1: map[string]int{
			"karam": 9711904500,
			"punam": 9711904343,
		},

		slice1: []string{
			"son",
			"daughter",
		},
	}

	fmt.Println(s.place)
	for i, v := range s.map1 {
		fmt.Println(i, v)
	}
	for i, v := range s.slice1 {
		fmt.Println(i, v)
	}

	//1. declare slice and map before func and use it in main
	//

}
