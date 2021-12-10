package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	// func literal aka anonymous self-executing func
	go func (){
		channel <- 1
	}()

	fmt.Printf("channel 1: %d\n",<-channel)

	// with buffer channel not ideal solution
	channel2 := make(chan int ,1)
	channel2 <- 2
	fmt.Printf("channel 2: %d\n",<-channel2)

}

