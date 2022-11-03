package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// encode the JSON

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

	err:= json.NewEncoder(os.Stdout).Encode(people)
	if err != nil{
		fmt.Println(err)
	}

	

}
