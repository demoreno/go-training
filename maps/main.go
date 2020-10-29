package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v \n", employees)

	fmt.Printf("No of pairs: %d \n", len(employees))

	fmt.Printf("The value for key: %q is %q \n ", "Dan", employees["Dan"])

	var accounts map[string]float64

	fmt.Printf("%#v \n", accounts["0x323"])

	// var m1 map[[5]int]string]

	people := map[string]float64{}

	people["David"] = 21.2
	people["Mary"] = 21.2
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 500,
		"EUR": 540,
	}
	fmt.Println(balances)

	v, ok := balances["RON"]

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Doesn't exists")
	}

	for k, v := range balances {
		fmt.Printf("key: %#v, value %#v \n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "X"}

	// fmt.Println(a == b)

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps equal")
	} else {
		fmt.Println("Not equal")
	}

	friends := map[string]int{
		"Dan":   40,
		"Maria": 25,
	}

	neighbors := friends

	friends["Dan"] = 50
	fmt.Println(neighbors)

	people1 := make(map[string]int)

	for k, v := range friends {
		people1[k] = v
	}

	people1["Anne"] = 18
	fmt.Println(people1)
	fmt.Println(friends)

	delete(friends, "Dan")
	fmt.Println(friends)
}
