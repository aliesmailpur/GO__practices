package main

import "fmt"

// create your own type "person" which will have an underlying
// type of "struct" so that it can store the following data:
//first name 
//last name 
// favorite ice cream flavors
// create two VALUES of TYPE person. print out the values.
// raging over the elements in the slice


type person struct{
	first_name string
	last_name string
	favorite_icecream []string
}


func main(){

	x:= person{
		first_name: "ali",
		last_name: "esmail pour",
		favorite_icecream: []string{"chocolate", "magnum", "milk icecream"},
	}


	fmt.Println(x.first_name)
	fmt.Println(x.last_name)
	fmt.Println(x.favorite_icecream)

	

	for i, v:= range x.favorite_icecream{
		fmt.Println(i, v)
	}
	




}