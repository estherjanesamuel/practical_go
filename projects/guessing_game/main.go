package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	const (
		MIN_VALUE = 1
		MAX_VALUE = 100
	)
	rand.Seed(time.Now().UnixNano())
	secret_number := rand.Intn(MAX_VALUE - MIN_VALUE) + MIN_VALUE
	fmt.Println("Guess the number!")
	for {
		
		// fmt.Println(secret_number)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please input your guess. ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.ParseInt(input, 10,64)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("You guessed: " ,guess)
		match, _ := cmp(int(guess), secret_number)
		
		switch match {
		case "Less":
			fmt.Println("Too Small!")
		case "Greater":
			fmt.Println("To Big!")
		case "Equal":
			fmt.Println("You Win!")
			return
		}
	}
}

func cmp(a,b int) (string,error)  {
	if a <  b {
		return "Less", nil
	}
	if a >  b {
		return "Greater", nil
	}
	return "Equal", nil
}