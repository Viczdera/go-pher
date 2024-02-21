/*
A race condition occurs when multiple threads  or in this case, go routines try to access and modify(Read or Write) the same data or memory address.
For example: if one thread tries to increase or add to an array and another thread tries to read it, this will cause a race condition.
Running the program with the race flag ("go run --race race.go") will display the race conditions
This example is illustrated below:
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Race conditions>")
	score := []int{0}
	waitG := &sync.WaitGroup{}

	waitG.Add(2) //two goroutines

	go func(wg *sync.WaitGroup) {
		fmt.Println("T One")
		score = append(score, 1)
		waitG.Done()
	}(waitG)
	go func(wg *sync.WaitGroup) {
		fmt.Println("T Two")
		score = append(score, 2)
		waitG.Done()
	}(waitG)

	//notifying main that there's a wait group
	waitG.Wait()

	fmt.Println(score)

}
