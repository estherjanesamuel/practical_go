package main

import (
	"fmt"
	"runtime"
	"time"
)


func f(from string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func p(till int, msg string)  {
	for i := 0; i < till; i++ {
		fmt.Println((i +1), msg)
	}
}

func main()  {
	
	runtime.GOMAXPROCS(4)
	fmt.Println()


	go func (msg string)  {
		fmt.Println(msg)
	}("going")
	
	go func ()  {
		fmt.Println("also going")
	}()
	
	go f("goroutine")

	go p(3,"ariefs")
	f("direct")

	time.Sleep(time.Second)
	fmt.Println("done")
}