package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	type book2 struct {
		title, autor string
		year         int
	}

	myBook := book{"The devine comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	bestBook := book{
		title:  "Animal Farm",
		author: "George Orwell",
		year:   1945,
	}

	fmt.Println(bestBook)

	lastBook := book{
		title: "Harry Potter",
	}

	fmt.Println(lastBook)

	lastBook.author = "JK Rowling"
	lastBook.year = 1991

	fmt.Println(lastBook)

	cBook := lastBook
	cBook.year = 2020

	fmt.Println(lastBook == cBook)

	fmt.Println(cBook)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "David",
		lastName:  "Moreno",
		age:       29,
	}

	fmt.Println(diana)

}
