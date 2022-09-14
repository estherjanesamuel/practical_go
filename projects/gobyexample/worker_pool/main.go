package main

import (
	"fmt"
	"time"
)

/*
	# In this example we'll look at how to implement
	# a woker pool using go routines and channels
*/
const numJobs = 1000

func main()  {
	/*
		# in order to use our pool of workers we need to **send them work**
		# and **collect their results**. we make 2 channels for this
	*/
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This startup 3 workers, initially blocked because these are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// here we send 5 `jobs` and then `close` that channel to indicate that's all the work we have
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	/*
		# finally we collect all the results of the work. this also ensure that the worker 
		# goroutines have finished. an alternatives way to wait for multiple go routines 
		# is to use a waitgroup
	*/
	for a := 1; a <= numJobs; a++ {
		<- results
	}
}

/*
	# here's the worker, of which we'll run several concurrent instances.
	# these worker will receive work on the `jobs` channel and send the corresponding
	# results on `results`. we'll sleep a second per job to simulate an expensive task.
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <-j * 2
	}
}