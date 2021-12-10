package main

import "fmt"


func main() {

	// channel diractions for thread safety
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	// send
	go send(eve,odd,quit)
	// receive 
    receive(eve,odd,quit)

	fmt.Println("about to exit")
}

// send
func send(eve chan<- int,odd chan<- int,quit chan<- int) {
	for i:= 0;i<100; i++{
		if i%2==0{
			eve <- i
		}else{
			odd <- i
		}
	}

	close(eve)
	close(odd)
	quit<-0
	close(quit)
}

// receive
func receive(eve <-chan int,odd <-chan int,quit <-chan int) {
	for {
		select {
		case v,ok:= <- eve:
			fmt.Println("from the eve channel:",eve,"the number:",v,"closed eve",ok)
		case v,ok:= <- odd:
			fmt.Println("from the odd channel:",odd,"the number:",v,"closed odd",ok)
		case v,ok:=<- quit:
			fmt.Println("from the quit channel:",<-quit,v,ok)
			return 
		}
	}
}



