package main

import (
	"fmt"
	"sync"
)

var x int64
var g sync.WaitGroup
var lock sync.Mutex

func add() {
	defer g.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func main() {
	g.Add(2)
	go add()
	go add()
	g.Wait()
	fmt.Println(x)
}
