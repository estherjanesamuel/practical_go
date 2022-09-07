package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	start := time.Now()
	c := Container{
		/*
			Note that the zero value of a mutex is usable as-is, 
			so no initialization is required here
		*/
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// this function increments a named counter in a loop
	doInc := func (name string, n int)  {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	/*
		run several goroutines concurrently; note that they all
		access the same `Container`, and two of them acces the same counter.
	*/ 
	wg.Add(3)
	go doInc("a", 10000)
	go doInc("a", 100)
	go doInc("b", 100)

	// wait a for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
	fmt.Println("elapsed time:", time.Since(start))
}

/*
	Container holds a map of counters; since we want to update
	it concurrently from multiple goroutines, we add a `Mutex` to synchronize access.
	Note that mutexes must not be copied, so if this `struct` is
	pass around, it should be done by pointer.
*/
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	/*
		Lock the mutex before accessing `counters`; unlock
		it at the end of the function using a defer statement
	*/
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}