package main

import (
	"fmt"
)

func main() {
	const value = 213
	var (
		i         = value
		f float64 = value
	)

	fmt.Println(i, f)
}
