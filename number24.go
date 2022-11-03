package main

import (
	"encoding/json"
	"fmt"
)

// marshal the user to JSON

type person struct{
	First string
	Last string
	Age int
}

func main(){

	p1:= person{
		First : "ali",
		Last: "tehrani",
		Age: 21,
	}
	p2:= person{
		First: "james",
		Last: "bond",
		Age: 50,
	}

	var people []person

	people = []person{p1,p2}

	bs, err:= json.Marshal(people) 
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))




}