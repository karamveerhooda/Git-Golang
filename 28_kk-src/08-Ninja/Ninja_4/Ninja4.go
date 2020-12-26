/*Starting with this code, sort the []int and []string for each person.
solution: https://play.golang.org/p/jz_llY1aPp
video: 128
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

}
