package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	n      int64
	wgr    sync.WaitGroup
	rwlock sync.RWMutex
)

func writer() {
	rwlock.Lock()
	n++
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wgr.Done()
}

func reader() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wgr.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wgr.Add(1)
		go writer()
	}

	for i := 0; i < 1000; i++ {
		wgr.Add(1)
		go reader()
	}
	wgr.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(n)
}
