package main

import "fmt"

// create a slice to store the names of all of the states in the united states of America.use make 
// and append to do this.Goal:do not have the array that underlines the slice created moore than once.
// what is the length of your slice? what is the capacity? Print out all of the values, along
// with their index position, without using the range clause. 



func main(){
	x:= make([]string, 50, 50)
	fmt.Println(len(x))
	states:= []string{"Alabama", "Alaska", "Arizona","Arkansas","California", "Colorado","Connecticut","delaware",
    "Florida", "Georgia","Hawaii","Idaho","illinois","indiona","lowa","kansass","kentucky","Louisiana","maine",
	"maine","maryland","massachusetts","michigan","minnesota","mississippi","missouri","montana","nebreska","nevada",
	"new hampshire","new jersey","new mexico"," new york","north carolina","north dakota","ochio","oklahoma","oregon",
	"pennsylvania","rhode island","south california","south dakota","tennessee","texas","utah","vermont","virginia",
	"washington","west virginia","wisconsin","wyoming"}
	for i,v:= range states{
		fmt.Println(i,v)
	}
	
}