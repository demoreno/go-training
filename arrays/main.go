package main

import "fmt"

func main() {
	// var numbers [4]int

	var a1 = [4]float64{}
	fmt.Printf("%#v \n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%v \n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "Jhon"}
	fmt.Printf("%v \n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%v \n", a4)

	a5 := [...]int{1, 2, 3, 4, 5, -10, 60}
	fmt.Printf("Length is %d \n", len(a5))
	fmt.Printf("%v \n", a5)

	a6 := [...]int{1,
		2,
		3,
		4,
		5,
	}
	fmt.Printf("%v \n", a6)

}
