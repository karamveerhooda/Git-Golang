/*Hands-on exercise #5
Starting with this code, sort the []user by
age
last
Also sort each []string “Sayings” for each user
print everything in a way that is pleasant
solution: https://play.golang.org/p/8RKkdtLl6w */
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First      string
	Last       string
	Age        int
	Sayings    []string
	MoreSaying []string
}
type byage []user
type bylast []user

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
		MoreSaying: []string{
			"let's",
			"sort",
			"this",
			"string",
			"also",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
		MoreSaying: []string{
			"again",
			"try",
			"till",
			"you not",
			"get succeed",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
		MoreSaying: []string{
			"Don't",
			"stop in between",
			"otherwise stoping",
			"become habit",
			"without reaching the goal",
		},
	}

	//James bond 32
	//				"Shaken, not stirred",
	//				"Youth is no guarantee of innovation",
	//				"In his majesty's royal service",

	users := []user{u1, u2, u3}

	for _, v := range users {
		fmt.Println(v.First, " ", v.Last, " ", v.Age)

		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {

			fmt.Println("\t", val)
		}
		sort.Strings(v.MoreSaying)
		fmt.Println("Moresaying")
		for _, moreval := range v.MoreSaying {
			fmt.Println("\t", moreval)
		}
	}
	//sort by age
	fmt.Println("--------------------Sort By Age----------------------------------")
	sort.Sort(byage(users))
	for _, v := range users {
		fmt.Println(v.First, " ", v.Last, " ", v.Age)

		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {

			fmt.Println("\t", val)
		}
		sort.Strings(v.MoreSaying)
		fmt.Println("Moresaying")
		for _, moreval := range v.MoreSaying {
			fmt.Println("\t", moreval)
		}
	}
	fmt.Println("--------------------Sort By Last----------------------------------")
	//sort by Last
	sort.Sort(bylast(users))
	for _, v := range users {
		fmt.Println(v.First, " ", v.Last, " ", v.Age)

		sort.Strings(v.Sayings)
		for _, val := range v.Sayings {

			fmt.Println("\t", val)
		}
		sort.Strings(v.MoreSaying)
		fmt.Println("    Moresaying")
		for _, moreval := range v.MoreSaying {
			fmt.Println("\t", moreval)
		}
	}

}

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (l bylast) Len() int           { return len(l) }
func (l bylast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l bylast) Less(i, j int) bool { return l[i].Last < l[j].Last }

// your code goes here
