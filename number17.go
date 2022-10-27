package main

import "fmt"

// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in

func main(){
	xi:= []int{1,2,3,4,5,6,7,8,9,}
	s:=foo(xi...)
	fmt.Println(s)
	x2:= []int{10,11,12,13,14,15,16,17,18,19}
	r:=bar(x2)
	fmt.Println(r)
	
	
	
	
}


func foo(x ...int) int{
	total:= 0
	for _, v:= range x{
		total +=v
	}
	return total
	 
}

func bar(b []int) int{
	total :=0
	for _, v:= range b{
		total += v
	}
	return total
}

