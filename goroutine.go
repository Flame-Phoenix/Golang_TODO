package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter = 0

func increament() {
	counter++
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
			// this printf shows the race condition
			fmt.Println(counter)
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter is: %d\n", counter)
}
