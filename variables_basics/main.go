package main

import "fmt"

func main() {
	var age int = 29
	fmt.Println("Age", age)

	var name = "David"
	fmt.Println("Name is", name)

	s := "Learning Go"
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2020
	fmt.Println(car, year)

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	// _, _ = i, j //mute error

	j, i = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
