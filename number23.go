package  main

import "fmt"

// create a person struct
// create a func called "challengeme" with a *person as a parameter
// change the value stored at the *person address
// create a value of type person
// print out the value
// in func main
// call changeme
// in func main
// print out the value

type person struct{
	name string
}

func challengeme(s *person) {

	s.name = "esi"


	

	

}


func main(){
	p1:= person{
		name: "ali",
	}
	fmt.Println(p1)

    challengeme(&p1)
	
	fmt.Println(p1)

	
}