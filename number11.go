package main

import "fmt"

// using the code from the previous example
// use SLICING to create the following slices
// [42,43,44,45,46]
// [47,48,49,50,51]
// [43,44,45,46,47]


func main(){
	x:= []int{42,43,44,45,46,47,48,49,50,51}

	x = append(x[:5] )
	fmt.Println(x)
	y :=append(x[5:10] )
	fmt.Println(y)
	a:=append(x[2:7])
	fmt.Println(a)
	b:=append(x[1:6])
	fmt.Println(b)
}