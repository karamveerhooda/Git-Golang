package main

import "fmt"

func main() {

	//1. using composite literal make an array , assign value and print it
	arr1 := [5]int{}

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr1)

	//2. Using  composite literal, make a slice , assign a value and print it
	slice1 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	for i, v := range slice1 {
		fmt.Println(i, v)

		//3. Take out slice of slice from above slice

		fmt.Println(slice1[:2])
		fmt.Println(slice1[2:])
	}
	fmt.Println(slice1[3:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[2:6])

	//4. append operation in slice

	slice2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice2 = append(slice2, 52)
	fmt.Println(slice2)
	slice2 = append(slice2, 53, 54, 55)
	fmt.Println(slice2)

	x := []int{56, 57, 58, 59, 60}
	slice2 = append(slice2, x...) // this .... is called variadic type
	fmt.Println(slice2)

	//5. Deleting from the slice

	//slice3 := append(slice2[:2])

	//slice4 := append(slice2[5:])

	//slice3 = append(slice3, slice4...)

	fmt.Println(append(slice2[:2], slice2[5:]...))

	// 6: Create a slice to store the names of all of the states in the United States of America.
	//What is the capacity?
	//Print out all of the values, along with their index position in the slice, without using the range clause.

	usastatescount := make([]string, 50, 50)
	usastatescount = []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`,
		` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`,
		` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Println(len(usastatescount))
	fmt.Println(cap(usastatescount))

	for i := 0; i < len(usastatescount); i++ {
		fmt.Println(i, usastatescount[i])
	}

	/*7. Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	  "James", "Bond", "Shaken, not stirred"
	  "Miss", "Moneypenny", "Helloooooo, James."

	  Range over the records, then range over the data in each record.
	  solution: https://play.golang.org/p/FMM4c2PhZg */
	sliceofslice1 := []string{"James", "Bond", "Shaken, not stirred"}
	sliceofslice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	sliceofslice := [][]string{sliceofslice1, sliceofslice2}
	for i, xs := range sliceofslice {

		fmt.Println("index iteration :", i)
		for j, val := range xs {
			fmt.Println(j, val)
		}

		fmt.Println()

	}
	/*Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

		solution: https://play.golang.org/p/nTzSlRa9_A

	*/

	map1 := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunses`},
	}

	fmt.Println(map1)
	//	fmt.Println("")

	for k, v := range map1 {
		fmt.Println("for Key:", k)
		for v, realvalue := range v {
			fmt.Println("\t\t", v, realvalue)
		}
	}

	/*9. Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
	 */
	fmt.Println()

	//mymap["Australia"] = 61
	map1["sterling"] = []string{"james", "bond", "heera"}
	for k, v := range map1 {
		fmt.Println("for Key:", k)
		for v, realvalue := range v {
			fmt.Println("\t\t", v, realvalue)
			//What is the length of your slice?
		}
	}

	fmt.Println("deleting from map")
	delete(map1, "sterling")
	for k, v := range map1 {
		fmt.Println("for Key:", k)
		for v, realvalue := range v {
			fmt.Println("\t\t", v, realvalue)
		}
	}

	mymap2 := map[string][]string{
		`hooda`:   []string{`karam`, `kk`},
		`khokhar`: []string{`punam`, `Ravi`},
		`joon`:    []string{`rajiv`, `pushkar`},
	}

	for k, val := range mymap2 {
		fmt.Println("for keys in map :", k)
		for val, innervalue := range val {
			{
				fmt.Println(val, innervalue)
			}
		}

	}

	/*Declaration of various Grouping data element
	1. Array  -- arr1 := [5]int{}
	2. Slice - composite literal --  slice1 := []int{}
				for i,v := range arr{
					print(i,v)
				}
	3. Slice - for range  -- for i,v := range slice1{print(i,v)
	4. Slice - Slicing a slice  -- slice1(start index:end index)
	5. Append to slice -- append(slice) -- append slice1(slice1,value)
	6. Slice - make , make([]T,length,capacity)


	7. Slice - multi-dimensional slice --  [][]slice1{[]T,[]T}
							for i,v := range slice1{
								for v,realvalue := range v{
									print v,realvalue
								}
							}
	8. Map   -- map1 := map[string][]string
	9. Map - add, range, delete  --

	*/

}
