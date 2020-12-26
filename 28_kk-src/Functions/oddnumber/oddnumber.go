package main

import "fmt"

func main() {

	oddslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(oddslice)

	x := odd(sum, oddslice...) /// 1 line

	fmt.Println("sum of odd numbers :", x)

}

func sum(y ...int) int { ///2 line

	var sum int

	for _, v := range y {
		if v%2 == 1 {
			//	newodd = append(newodd, v)
			sum += v
		}
	}
	return sum
}

func odd(f func(x ...int) int, vi ...int) int { ////3rd line

	var newodd []int
	//var sum int

	for _, v := range vi {
		if v%2 == 1 {
			newodd = append(newodd, v)
			//	sum += v
		}
	}
	fmt.Println(newodd)
	return f(newodd...) //// 4th line.
}
