package main

import "fmt"

// create a func which "encloses" the scope of a variable:

func ali() func () int{
	var x int
	return func ()int {
		x++
		return x
	}
}



func main(){

	a:= ali()
	b:= ali()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}