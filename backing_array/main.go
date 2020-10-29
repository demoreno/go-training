package main

import "fmt"

func main(){
	/*s1 := []int{10,20,30,40,50}
	s3, s4 := s1[0:2],s1[1:3]

	s3[1] = 600
	fmt.Println(s1) 
	fmt.Println(s4) 

	arr1 := [4]int{10,20,30,40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)*/

	cars := []string{"Ford", "Honda", "Toyota", "Chery"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"

	fmt.Println(cars, newCars)

	s1 := []int{10,20,30,40,50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice),cap(newSlice))

	newSlice = s1[2:5]  
	fmt.Println(len(newSlice),cap(newSlice))

	fmt.Printf("%p \n", &s1)

	fmt.Printf("%p %p \n", &s1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("s1 :", s1)

}