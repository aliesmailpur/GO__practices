package main

import (
	"fmt"
	"math"
)

// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create area p r 2
// square are L * W
// create a type shape which defines and interface
// as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct{
	length float64
}
type circle struct{
	radius float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64{
	return s.length * s.length
}

type shape interface{
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
	
}

func main(){
	circ := circle{
		radius: 12.456,
	}

	squa:= square{
		length: 7.987,
	}
	info(circ)
	info(squa)
}
