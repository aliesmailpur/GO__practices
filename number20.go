package main

import "fmt"

// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

func ali() func() int{
	return func() int{

		return 541

	}
}


func main(){

	s:= ali()
	x:= s()
	fmt.Println(x)



}