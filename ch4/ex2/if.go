package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}

	for _, num := range nums {
		switch {
		case num%6 == 0:
			fmt.Println("Six!")
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
