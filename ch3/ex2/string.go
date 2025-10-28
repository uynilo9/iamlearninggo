package main

import (
	"fmt"
)

func main() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	char := string(runes[3])

	fmt.Println(char)
}
