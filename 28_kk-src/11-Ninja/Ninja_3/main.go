/*Hands-on exercise #3
Create a struct “customErr” which implements
 the builtin.error interface. Create a func “foo” that has a value of type
 error as a parameter. Create a value of type “customErr” and pass it into “foo”.
 If you need a hint, here is one.
solution:
	https://play.golang.org/p/ixeowY2fd2
assertion
		https://play.golang.org/p/pbl2kCYsM0
conversion
		https://play.golang.org/p/1ldiBdkdzA
video: 178
7. Interface type:Interface type is any other normal type like int. In example Ninja11->Ninja3
        func (ce customeErr) Error() string {
            	return "wtf"
               }
        Here we are using builtin.Error() (builtin is package and Error() is the only method of interface )
        And here error is the name of Interface, below is the golang snippet of error interface:
        type error interface {
           Error() string
             }
        So if anyone want to call implement error interface, he have to implement its only method Error()
*/
package main

import "fmt"

type customeErr struct {
	name string
	item string
}

func (ce customeErr) Error() string {
	return "we are here really"
}
func foo(err error) {
	fmt.Println(err)
}
func main() {
	c1 := customeErr{
		name: "karam",
		item: "phone",
	}

	fmt.Println(c1.Error())
	foo(c1)
}
