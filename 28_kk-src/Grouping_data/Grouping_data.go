package main

import "fmt"

func main() {
	//how to read
	x := []int{22, 33, 44, 55, 66}
	//.......slices of int and there are the values
	//SLICE allows you to group together value of the same type
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
	//slicing the slice means to take out any specific position element out of slice or to print specific range element from slice

	y := []int{12, 13, 14, 15, 23, 24}
	fmt.Println(y)
	fmt.Println(y[:])
	fmt.Println(y[0:1]) // This is called a slicing a slice process. video 80 of Todd mcleod
	fmt.Println(y[1:2])
	fmt.Println(y[2:3])
	fmt.Println(y[3:4])
	fmt.Println(y[4:5])
	fmt.Println(y[5:6])

	for x := 0; x <= 5; x++ {
		fmt.Println(x, y[x])
		fmt.Println()
	}

	// To append value in slice
	// func append(slice []T,elements ...T) []T  ---- here ...T is called variadic parameters

	y = []int{4, 26, 27, 28}

	x = append(x, 111, 121, 131)
	fmt.Println(x)
	fmt.Println("next line")
	z := append(x, y...)

	fmt.Println(z)

	//to delete from slice

	z = append(z[:2], z[4:]...)
	fmt.Println(z)

	// How to make use of make -- make(slice of T[], length ,capacity)

	newmake := make([]int, 10, 12)
	fmt.Println(newmake)
	fmt.Println(len(newmake))
	fmt.Println(cap(newmake))

	newmake[1] = 1
	newmake[9] = 10
	newmake = append(newmake, 11)
	fmt.Println(newmake)
	fmt.Println(len(newmake))
	fmt.Println(cap(newmake))
	newmake = append(newmake, 12)
	newmake = append(newmake, 13)
	newmake = append(newmake, 14)
	newmake = append(newmake, 15)
	newmake = append(newmake, 16)
	newmake = append(newmake, 17)
	newmake = append(newmake, 18)
	newmake = append(newmake, 19)
	newmake = append(newmake, 20)
	newmake = append(newmake, 21)
	newmake = append(newmake, 22)
	newmake = append(newmake, 23)
	newmake = append(newmake, 24)
	newmake = append(newmake, 25)
	fmt.Println(newmake)
	fmt.Println(len(newmake))
	fmt.Println(cap(newmake))

	multidimesionalslice := [][]int{newmake, z}

	fmt.Println(multidimesionalslice)

	// How to use maps:
	m := map[string]int{"karam": 11,
		"punam": 12,
	}
	fmt.Println(m)
	fmt.Println(m["karam"])

	v, ok := m["tejus"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["karam"]; ok {
		fmt.Println("ya ok", v)
	}

	mymap := map[string]int{"India": 91, "USA": 1, "UK": 44}

	fmt.Println(mymap)

	v, ok = mymap["USA"]
	fmt.Println(v)

	if v, ok = mymap["Pak"]; ok {
		fmt.Println("This is incorrect country", v)
	} else if v, ok = mymap["India"]; ok {
		fmt.Println("country is India", v)
	}

	// how to add the key in maps

	mymap["Australia"] = 61

	for k, v := range mymap {

		fmt.Println(k, v)
	}

	// How to delete entry from a map

	delete(mymap, "Australia")
	fmt.Printf("deleting the element from mymap\n")
	for k, v := range mymap {
		fmt.Println(k, v)
	}

}
