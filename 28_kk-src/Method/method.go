package main
import "fmt"

type ingred struct{
	Greenvegie string
	vegie2 string
}

type receipe struct{
	ingred
	tomato bool
}

//func (r receiver) identifer(parameter) (return(s)) {...}
func (r receipe)sabjireciepe(parameterfrommain string) string {

	fmt.Println("in sabjireciepe")
	fmt.Println(r.ingred,r.tomato,r.ingred.vegie2)
	fmt.Println(parameterfrommain)
	//var message string = "Sarsoo da saag is my first receipe"
	//fmt.Sprint(message, receipe1.ingred, receipe1.tomato)
	return "done"
}
func (ing ingred)vegrecieper(){
	fmt.Println(ing.vegie2,ing.Greenvegie)
}


func main() { 

	receipe:= receipe{
		ingred: ingred{
			"palak",
			"Raddish",
		},
		tomato: true,
	}

	fmt.Println("in main", receipe)

	fmt.Println(receipe.sabjireciepe(receipe.ingred.vegie2))
	receipe.vegrecieper()
}