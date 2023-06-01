package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func goRoutines2() {
	fmt.Println("Phase 13, Let's go!")

	msgVar := "Hello Part 2 of GoRoutines"
	// using WaitGroup to wait for groups of Go Routines to complete
	wg.Add(1)
	go func(s string) {
		fmt.Println(s)
		wg.Done()
	}(msgVar)

	wg.Wait()

	// sync.Mutex and sync.RWMutex to protect data access of data that is going to be accessed in multiple locations in the application

	// by-default, Go will use CPU THREADS EQUAL TO THE AVAILABLE CORES
	// we can change it with runtime.GOMAXPROCS(enter number of threads)
	// MORE THREADS can INCREASE PERFORMANCE, BUT TOO MANY CAN SLOW IT DOWN

	// Don't create Go routines in libraries
	// when creating Go Routines, know how will they end
}
