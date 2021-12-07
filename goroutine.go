package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter = 0
var mu sync.Mutex

func increament() {
	// mutex fix the race condition
	mu.Lock()
    defer mu.Unlock()
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
			runtime.Gosched()
			increament()
			
		}()
	}
	wg.Wait()
	fmt.Printf("Counter is: %d\n", counter)
}
