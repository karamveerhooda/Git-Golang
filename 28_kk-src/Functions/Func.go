/*Different func module
1. normal func
2. Variadic func
3 unfurling of slice
4. defer of func (defer means delaying the execution of function till all other function get executed)
<https://play.golang.org/p/F3LN16BmfEI>
*/
/*5. Methods - in this we use the receiver as a data type with function, see below line
func (r receiver) identifier(parameters) (return(s)){...}.
6. Function as a expression
7. function return type, means function returning a function type. As function are normal return types like Int in GOLANG

Q: how to read  (r reciever ) identifier(parameter)  --
Identfier(parameter) func have access to the value of type receiver
AGAIN SAY : A func have access to the value(r) of receiver type.
https://play.golang.org/p/JnjLU9enrW5

in above code
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	reciepe()
	/*func calling semantics
	func (r receiver) identifier (paramters)(retrun (s){...code})
	-----(called paramter attached to the type)
		Q :now the difference between paramter and argument?
		A: When calling function like
		// In GO everything is call by value only that's it
	*/
	// firstreciepe("First reciepe")
	// ingred1, ingred2 := secondreciepe("Ans : sarso ", "palak")
	// fmt.Println(ingred1)
	// fmt.Println(ingred2)
	fmt.Println(secondreciepe("Ans : sarso ", "palak"))
	// // Variadic paramter func
	var mealgood bool
	var newpricelist []int
	xi := []int{5, 10, 15, 20}
	sum, mealgood, newpricelist := pricerangeeofmeal(xi...) //xi...is called unfurling of slice
	// here pricerangeeofmeal is looking for int but xi is slice of int, whereas (x ...int) is expecting
	//unlimited number of int. So we need to unfurl the slice like (x ...int) where we are calling the
	//pricerangeeofmeal like in above line 	sum, mealgood, newpricelist := pricerangeeofmeal(xi...)
	pricemenu := []string{"veg", "non-veg"}

	pricemenufunc("myprice")
	pricemenufunc("veg:", "non-veg", "jains")

	fmt.Println(pricemenu)
	fmt.Println()

	fmt.Println(sum, mealgood)
	for i, v := range newpricelist {
		fmt.Println(i, v)
	}

	//unfurling of array

	//Assigning a function to a variable, called function expression:

	x := func(y int) {
		fmt.Println("Assigning a function to a variable", y)
	}
	x(1980)
	// Func returning a func type. SO declare a func with func return type.
	// Function are first class citizen :-) --
	//1. function can be return type
	//2. function can be assigned to a variable
	//3. function can be passed as an argument to a function

	fmt.Printf("%T\n", bar())
	fmt.Printf("%T", bar()())
	fmt.Println("\nz :", bar()())

	/* SO CALLBACK is pass a func as an argument.
	Prog: create a function that sums even numbers from slice*/

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := sum(ii...)
	fmt.Println(s)

}

func sum(xii ...int) int {
	//evenlist []int
	s := 0
	var ii []int
	for _, v := range xii {
		if v%2 == 0 {
			s += v
			ii = append(ii, v)
		}
	}
	fmt.Println(ii)
	return s
}

func bar() func() int {
	return func() int {
		return 111
	}

}
func pricemenufunc(s string, x ...string) int {

	sums := 0

	fmt.Println(s)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	return sums
}
func pricerangeeofmeal(x ...int) (int, bool, []int) {
	var sum int
	for i, v := range x {
		fmt.Println(i, v)
		sum += v
	}

	var newpricelist []int = append(x, 50)

	fmt.Println("Sum of all meals :", sum)
	return sum, true, newpricelist
}

func reciepe() {
	fmt.Println("I am in reciepe")
}

func firstreciepe(a string) {
	fmt.Println("my", a)
}

func secondreciepe(firstingred string, secondingred string) (string, bool) {
	a := fmt.Sprint(firstingred, secondingred, " the two main ingredient of sarsoo da saag")
	b := true
	return a, b
}
