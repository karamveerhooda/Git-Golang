/*How to use interface --  Interface where we give the definition of method but not impement them there. 
						or in other words "Interface also called the collection of methods signature" 
						
Below program find out the various types that we have used as a struct using interface tasterating
--we find using interface taste have one method tasterating. 
-- Tasterating method implemented with two types one is ing ingred and another r receipe
-- func sabjireciepe check the interface t of type taste. 
-- So when we pass the struct of receipe1 or receipe2 of receipe type, sabjireciepe switch statement show the respective type */

package main
import (
	"fmt"
//	"go/types"
)
type ingred struct{
	Greenvegie string
	vegie2 string
}

type receipe struct{
	ingred
	tomato bool
}

type taste interface{
	tasterating()
	tasterRatingComment(comment string)
}

func (ing ingred) vegrecieper(){
	fmt.Println(ing.vegie2,ing.Greenvegie)
}
//method using interface 
func (r receipe) tasterating(){
	fmt.Println("I am in tasterating of type receipe", r.tomato)
}
//another method using interface
func(ing ingred) tasterating(){
	fmt.Println("I am in tasterating method with type ingred",ing.vegie2)
}
//method using interface 
func (r receipe) tasterRatingComment(comment string){
	fmt.Println("The comment receives on this reciepe is :", comment )
}
//another method using interface
func(ing ingred) tasterRatingComment(comment string){
	fmt.Println("The comment receives on this reciepe is :", comment)
}


//func (r receiver) identifer(parameter) (return(s)) {...}
func sabjireciepe(t taste)  {
	
	switch t.(type){

	case receipe:
		fmt.Println("case true for receipe type receipe", t.(receipe).tomato)
	case ingred:
		fmt.Println("case true for ingred type ingred", t.(ingred).vegie2)
	}

	fmt.Println("I am in sabjirecipe function", t)
	//	fmt.Printf("%/T",t.vegrecieper)
	}
	
func main() { 

	receipe1:= receipe{
		ingred: ingred{
			"palak",
			"Raddish",
		},
		tomato: true,
	}

	receipe2 := receipe{
		ingred: ingred{
			"sarso",
			"methi",
		},
		tomato: false,
	}

	receipe1.tasterating()
	fmt.Println("-------")
	receipe2.tasterating()
	fmt.Println("-------")

	//fmt.Println(receipe2.tasterating())
	receipe1.tasterRatingComment("awesome")
	fmt.Println("-------")
	receipe2.tasterRatingComment("ok not so good")
	
	fmt.Println("-------")
	fmt.Println("-------")
	sabjireciepe(receipe1)
	fmt.Println("-------")
	sabjireciepe(receipe2)
	fmt.Println("-------")
	sabjireciepe(receipe2.ingred)
}