package main

import "fmt"

func main() {
	const golang = "golang"

	runes := []rune(golang)
	for i := 1; i <= len(runes); i++ {
		fmt.Printf("%c\n", runes[:i])
	}
	for i := len(runes) - 1; i >= 1; i-- {
		fmt.Printf("%c\n", runes[:i])
	}
}
