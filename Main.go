package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOMAXPROCS(8))
	for {
		fmt.Println(runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}
