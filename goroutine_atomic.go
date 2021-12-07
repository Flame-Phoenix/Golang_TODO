package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
func increment() {
	// this printf shows the race condition
}

func main() {

	var wg sync.WaitGroup
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// atom is faster but its not doing it in syncronus order
			atomic.AddInt64(&counter,1)
			// that print shows that
			fmt.Println(atomic.LoadInt64(&counter))
			defer wg.Done()
		}()
	}
	wg.Wait()
	// the result is right tho
	fmt.Printf("Counter is: %d\n", counter)
}
