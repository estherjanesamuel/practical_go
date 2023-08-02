package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// listen for the termination signal
		stop := make(chan os.Signal,1)
		// we need to reserve the buffer size to 1
		// so the notifier are not blocked

		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// block until termination signal recive
		<- stop

		// essentialy the cancel() is broadcasted to all
		// go routines that call .DOne()
		// the return context's Doe channel is closed when 
		// the returned cancel functun is called
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				fmt.Println("Break the first loop")
				return
			case <- time.After(1 * time.Second):
				fmt.Println("Hello from the first loop")
			}
		}
	}()

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for {
			select {
			case <- ctx.Done():
				fmt.Println("Break the second loop")
				return
			case <- time.After(1 * time.Second):
				fmt.Println("Hello from the second loop")
			}
		}
	}()

	// wait for ongoing process to finish
	wg.Wait()
	fmt.Println("Main done")
}