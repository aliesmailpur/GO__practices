package main

import "fmt"

// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assin 10 VALUES
// range over the slice and print the values out
// using format printing 
// print out the TYPE of the slice



func main(){
	x:= []int{42,43,44,45,46,47,48,49,50,51}

	for i, v:= range x{
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n", x)
}