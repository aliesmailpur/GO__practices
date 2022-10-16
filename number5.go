package main

import "fmt"

// Using iota, create 4 constants for the last 4 years. Ptint the 
// constant values.

const (

	a = 2018 + iota
	b = 2018 + iota
	c = 2018 + iota
	d = 2018 + iota

)

func main(){

	fmt.Println(a, b, c , d)


}