package main

import (
	"fmt"
	"runtime"
)

func main() {
	// GOOS : operating system
	// GOARCH : 
	fmt.Println(runtime.GOOS,runtime.GOARCH)
}
