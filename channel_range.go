package main

import "fmt"

// send
func ping(pings chan<- int) {
    for i:= 0; i<5 ; i++{
		pings <- i
	}
	close(pings)
}


func main() {

	pings := make(chan int)
	// send
	go ping(pings)
	// receive 
    for v:= range pings {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}