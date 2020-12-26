package main

import (
	"fmt"
)

var map1 map[string]int
var map2 map[string][]string

type mystruct struct {
	exam          string
	subject       []string
	candidatename map[string]int
}

func main() {
	map1 := map[string]int{
		"karamveer": 111,
		"punam":     222,
	}

	for i, v := range map1 {
		fmt.Println(i, v)
	}

	map2 := map[string][]string{

		"firstname": {"karam", "punam", "rajiv"},
		"lastname":  {"hooda", "khokhar", "joon"},
	}

	for i, v := range map1 {
		fmt.Println(i, v)
	}
	for i, v := range map2 {
		fmt.Println(i)
		for v, val := range v {
			fmt.Println(v, val)
		}
	}

	mystruct := mystruct{
		exam:    "10th board",
		subject: []string{"eng", "hindi", "math", "science", "SST"},
		candidatename: map[string]int{
			"pushkar": 9711,
			"bhavika": 8711,
			"chinu":   1199,
		},
	}

	fmt.Println(mystruct.exam)
	for i, v := range mystruct.subject {
		fmt.Println(i, v)
	}
	for i, v := range mystruct.candidatename {
		fmt.Println(i, v)
	}

}
