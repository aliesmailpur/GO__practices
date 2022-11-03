package main

import (
	"encoding/json"
	"fmt"
)

// unmarshal the JSON for the previous exercise

type person struct{
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main(){
	s:=`[{"First":"ali","Last":"tehrani","Age":21},{"First":"james","Last":"bond","Age":50}]`

	var people []person
	
	err:=json.Unmarshal([]byte(s), &people) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)




}