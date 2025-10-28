package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	first := Employee{
		"Yat-sen",
		"Sun",
		1112,
	}

	second := Employee{
		firstName: "Kai-shek",
		lastName:  "Chiang",
		id:        1031,
	}

	var third Employee
	third.firstName = "Ren-Yu"
	third.lastName = "Lin"
	third.id = 213

	fmt.Println(first, second, third)
}
