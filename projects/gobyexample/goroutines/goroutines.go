package main

/* a goroutine is a lightweight thread of execution */

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
	going := "going"
	/*
		# suppose we have a function call `f(s)`. here's how 
		# we'd call that in that usual way, running it synchronously.
	*/
	f("direct")

	/*
		# to invoke this function in a goroutine, use `go` before the function call => `go f(s)`.
		# this new goroutine will execute concurrently with the calling one / this will kick's of the goroutine
	*/
	go f("goroutine")

	//you can also start a goroutine fo an anonymous function call 
	go func (msg string)  {
		fmt.Println(msg)
	}(going)
	// going here is a capture variable (closure)
	go func ()  {
		fmt.Println("also going")
	}()

	go p(3,"ariefs")

	/*
		# our two function calls are running asynchronously in 
		# seperate goroutines now.wait for them to finish
		# for a more robust approach, use a waitgroup
	*/ 
	time.Sleep(time.Second)
	fmt.Println("done")
}