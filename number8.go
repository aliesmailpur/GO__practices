package main

import "fmt"

// print out the remainder which is found for 
// each number between 10 and 100 when it is divided by 4

func main(){
	for i:=10;i<=100;i++{
		fmt.Printf("when %d is divided by 4 the remainder is %d\n" , i , i%4)
	}
}