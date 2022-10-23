package main 

import "fmt"

// create a program that uses a switch statement with no switch expression
// specified as a variable of TYPE string with the IDENTIFIER "favSport"



func main(){
	favSport:= "football"

	switch favSport{
	case "boxing":
		fmt.Println("go to boxing class")
	case "tennis":
		fmt.Println("go play tennis")
	case "football":
		fmt.Println("go play football")

	}


}