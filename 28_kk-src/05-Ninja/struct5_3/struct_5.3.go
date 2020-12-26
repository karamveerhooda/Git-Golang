package main

import "fmt"

type vehicle struct {
	doors string
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

/*type myanonymous struct {
	map1 map[string][]string
	cake []string
}*/

func main() {

	truck := truck{
		vehicle: vehicle{
			doors: "2 doors",
			color: "Blue",
		},

		fourwheel: true,
	}
	sedan := sedan{
		vehicle: vehicle{
			doors: "4 doors",
			color: "metallic grey",
		},

		luxury: true,
	}

	fmt.Println(truck)
	fmt.Println("----")
	fmt.Println(sedan)

	// 4. Anonymous Struct

}
