package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func increment() {
	for i := 0; i < 100000; i++ {
		value := counter
		value++
		counter = value
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// The provided Go code has a race condition on the counter variable. A race condition occurs when two or more goroutines access shared data concurrently, and at least one of them modifies the data. In this case, both go increment() goroutines are concurrently incrementing the counter variable without proper synchronization. This can lead to unpredictable and inconsistent results, as both goroutines may read and modify the counter simultaneously.