package main

import "fmt"

func main() {
	var a = 2
	var b = 3.2

	a = int(b)
	fmt.Println(a, b)

	// var x = 5
	// x = "5"

	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
