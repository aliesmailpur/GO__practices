package main

import "fmt"

// create a for loop using this syntax
// for condition {}
// have it print out the years you have been alive.

func main(){

	for i:=1990;i<=2022;i++{
		if i <= 2000{
			continue
		}

		if i >= 2000{
			fmt.Println(i)
		}
	}
}