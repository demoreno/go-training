package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%v \n", numbers)

	numbers[0] = 200
	fmt.Printf("%v \n", numbers)

	for i, v := range numbers {
		fmt.Println("index:", i, " value: ", v)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index, ", i, " value: ", numbers[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}

	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m

	fmt.Println("n is equal to m: ", n == m)

	m[0] = -1

	fmt.Println("n is equal to m: ", n == m)
}
