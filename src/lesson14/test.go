package main

import (
	"fmt"
	"runtime"
	"time"
)

func A() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func B() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	go A()
	go B()
	time.Sleep(time.Second)
}
