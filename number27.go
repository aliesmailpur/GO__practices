package main

import (
	"fmt"
	"sort"
)

// sort these codes out

func main(){
	xi:= []int{888,754,12,56,888,}
	xs:= []string{"a","z","h","d","s","v","p"}
	fmt.Println(xi,xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)


}