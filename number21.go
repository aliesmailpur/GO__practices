package main

import "fmt"

// for this exercise , pass a func into a func as an argument

func sum(x ...int) int{

	total := 0
	for _, v:= range x{
		total += v
	}

	return total



}

func odd(f func(xi ...int) int,  d ...int)int{
	var g []int
	for _,v:= range d{
		if v%2 !=0{
			g= append(g, v)
		}
	}
	return f(g...)

}


func main(){
	xi:= []int{1,2,3,4,5,6,7,8,9,}
	s:= sum(xi...)
	fmt.Println(s)

	c:= odd(sum, xi...)
	fmt.Println(c)

}

