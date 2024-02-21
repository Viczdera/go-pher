package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race conditions>")
	score := []int{0}
	waitG := &sync.WaitGroup{}
	mutX := &sync.RWMutex{}

	waitG.Add(3) //three goroutines

	go func(wg *sync.WaitGroup, mx *sync.RWMutex) {
		fmt.Println("One R")
		mx.Lock()
		score = append(score, 1)
		mx.Unlock()
		waitG.Done()
	}(waitG, mutX)
	go func(wg *sync.WaitGroup, mx *sync.RWMutex) {
		fmt.Println("Two R")
		mx.Lock()
		score = append(score, 2)
		mx.Unlock()
		waitG.Done()
	}(waitG, mutX)
	go func(wg *sync.WaitGroup, mx *sync.RWMutex) {
		fmt.Println("Three R")
		//	mx.Lock()
		//	score = append(score, 3)
		mx.RLock()
		fmt.Println(score)
		mx.RUnlock()
		waitG.Done()
	}(waitG, mutX)

	//notifying main that there's a wait group
	waitG.Wait()

	fmt.Println(score)

}
