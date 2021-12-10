package main

import (
	"fmt"
)

// send
func ping(pings chan<- string, msg string) {
    pings <- msg
}
// receive
func pong(pings <-chan string) {
	fmt.Println(<-pings)
}

func main() {

	// channel diractions for thread safety
	pings := make(chan string)
	// send
	go ping(pings, "passed message")
	// receive we remove go key word so the program block until it happens
    pong(pings)
    
	fmt.Println("about to exit")
}

// other way to do the above

// func ping(pings chan<- string, msg string) {
//     pings <- msg
// }

// func pong(pings <-chan string, pongs chan<- string) {
//     msg := <-pings
//     pongs <- msg
// }

// func main() {
//     pings := make(chan string, 1)
//     pongs := make(chan string, 1)
//     ping(pings, "passed message")
//     pong(pings, pongs)
//     fmt.Println(<-pongs)
// }