package main
import ("fmt")
//struct are composite aka aggregate aka complex data type
type student struct{
	name string
	rollno int
}
type coachingstudent struct{
	student
	attendance bool
}
func main() {

class11 := student{
	name: "karam",
	rollno: 33,
}
fiitjee := coachingstudent{
				student: student{
					name: "ajay",
					rollno: 11,
				},
				attendance: false,
}
fmt.Println(fiitjee)
	fmt.Println(fiitjee.student.name,fiitjee.student.rollno,fiitjee.attendance)
	fmt.Println(class11.name,class11.rollno)
	
}

/*
package main
import (
	"fmt"
)
type mystruct struct {
	first string
	last  string
}
func main() {
	mystruct1 := mystruct{
		first: "karam",
		last:  "hooda",
	}
	mystruct2 := mystruct{
		first: "punam",
		last:  "khokhar",
	}
	fmt.Println(mystruct1.first, mystruct1.last)
	fmt.Println(mystruct2.first, mystruct2.last)
}*/