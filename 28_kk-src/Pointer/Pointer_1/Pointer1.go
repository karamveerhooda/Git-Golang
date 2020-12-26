package main
import (
	"fmt"
)

/*1. Address of variable denoted by & for e.g. x address is &x  -- learn ampersand synonym to address
  2. Now *x i called dereferencing so it is readable as value at address x
  3. What is the type of address &x in point 1 -- it is *int read as pointer to an int , this is new type just like int , float or bool
  4. */
func main() {
	x := 5
	
	fmt.Println("Address of x is :",&x)
	fmt.Printf("The type of address is %T\n",&x)
/************************/	
// we can not assign &x to y int because &x is of *int type and y is int type. Both are of different type like you can not assign bool to an int
//	var y int = &x
//  but we can assign &x to a variable of *int(read as pointer to an int)
	var z *int = &x
	fmt.Println("Value of z is :",z)
	fmt.Printf("The type of address z is %T\n",z)

	z = &x
	fmt.Println("address of z is :",z)
	fmt.Println("The value stored at address of z is :", *z) // *z gives you the value stored at an address when you have the address. 
															//This *  is operator is called dereferencing means we are derferencing the address of z here.
															//when we use *int called pointer to an int or say pointer to where int is stored.
	fmt.Println("The value stored at x is :",*&x) // This is readable as dereferencing the value stored at address of x


	fmt.Println()
	foo(&x)
	fmt.Println("Address of x is :",&x)
	fmt.Printf("The type of x is %T\n",x)
	fmt.Println("Value stored at x is:",x)
}
func foo(y *int) {
	fmt.Println("I am after foo call y :",y)
	fmt.Println("Value of *y is :",*y)
	*y = 10
	fmt.Println("The address of y is :",&y)
	fmt.Println("Value stored at *y is:",*y)

	fmt.Printf("%T\n",y)
}