package main

import (
	"fmt"
)

func main() {
	var (
		greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
		s1        = greetings[:2]
		s2        = greetings[1:4]
		s3        = greetings[3:]
	)

	fmt.Println(greetings, s1, s2, s3)
}
