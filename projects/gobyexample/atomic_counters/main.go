package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	# The primary mechanismn for managing state in Go is
	# communication over channels. We saw this for example with
	# worker pools. There are a few other options for managing
	# state though. Here we'll look at using the `sync/atomic`
	# package for `atomic counters accessed by multiple goroutines.`
*/
func main()  {
	// we'll use an unsigned integer to represent our (always positive) counter
	var ops uint64

	// a waitgroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// we;ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func ()  {
			for c := 0; c < 1000; c++ {
				// to atomically increment the counter we use `AddUint64`,
				// giving it the memory address of our `ops` counter with the `&` syntax
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// wait until all the goroutines are done
	wg.Wait()

	/*
		# it's safe to access `ops` now because we know no other gorutine 
		# is writing to it. reading atomics safely while they are being updated is 
		# also possible, using functions like `atomic.LoadUint64`.
	*/

	fmt.Println("ops:", ops)
}