package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("Cities is equal to nil: ", cities == nil)
	fmt.Printf("Cities %#v \n", cities)
	fmt.Println(len(cities))

	// cities[0] = "London" is an error

	numbers := []int{2, 3, 4, 5}
	println("Numbers ", numbers)

	nums := make([]int, 2)
	fmt.Printf("%v \n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is", myFriend)

	friends[0] = "David"
	fmt.Println("My best friend is", friends[0])

	for index, value := range numbers {
		fmt.Printf("index %v value: %v \n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
