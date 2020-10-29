package main

import "fmt"

func main(){
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1,2,3}, []int{1,5,3}
	// fmt.Println(a == b)

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i]{
			eq = false
			break
		}
	}

	if len(a) != len(b){
		eq = false
	} 

	if eq{
		fmt.Println("equal", eq)
	} else {
		fmt.Println("not equal", eq)
	}
}