package main

import (
	"fmt"
	"sync"
	// "runtime"
)

var counter = 0
var mu sync.Mutex

func increment() {
	// using mutex to fix the race condition
	defer mu.Unlock()
	mu.Lock()
	counter++
	// this printf shows the race condition
	fmt.Println(counter)
}

func main() {

	var wg sync.WaitGroup
	gs :=100

	for i := 0; i < gs; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			// not make sense to use with mutex lock
			// runtime.Gosched()
			increment()
			
		}()
	}
	wg.Wait()
	fmt.Printf("Counter is: %d\n", counter)
}
