package main

import "fmt"

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.5
	const secondsInHour = 3600

	duration := 234
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// x, y := 5, 0

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	// if you don't specific the value of a constant their asign the previous value
	/* const (
		min1 = -500
		min2 =
		min3 =
	) */

	fmt.Println(min1, min2, min3)

	// Constants Rules
	const temp int = 100

	// You can't initiate  a constant at runtime
	// const power = Math.pow(2, 3)

	//

}
