package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	for i := 0; i < 30; i++ {
		time.Sleep(500 * time.Millisecond) // mimics work
		printProgressBar(i+1, 30, "Progress", "Complete", 25, "=")
	}
}

func printProgressBar(iteration, total int, prefix, suffix string, length int, fill string)  {
	percent := float64(iteration) / float64(total)
	filledLength := int(length * iteration / total)
	end := ">"

	if iteration == total {
		end = "="
	}

	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %f%% %s", prefix, bar, percent, suffix)
	if iteration == total {
		fmt.Println()
	}
}