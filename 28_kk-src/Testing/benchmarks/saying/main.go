package saying

import (
	"fmt"
)

//Greet func
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}

/*Here we are testing the saying/main program return correct string or not
Like when we pass James to Greet method of saying/main it should return
 Welcome my dear james.*/
