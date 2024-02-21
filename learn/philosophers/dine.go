package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChopS struct {
	sync.Mutex
}
type Phil struct {
	num             int
	leftCS, rightCS *ChopS
}

// eat method
func (p Phil) eat(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	//eats three times
	for i := 0; i < 3; i++ {
		ch <- struct{}{}
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d\n", p.num)
		fmt.Printf("finishing eating %d\n", p.num)
		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-ch
	}
}
func main() {
	cS := make([]*ChopS, 5)
	pH := make([]*Phil, 5)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		cS[i] = new(ChopS)
	}
	for i := 0; i < 5; i++ {
		lr := []int{i, (i + 1) % 5}
		randomOrder(lr)
		pH[i] = &Phil{
			num:     i + 1,
			leftCS:  cS[lr[0]],
			rightCS: cS[lr[1]],
		}
		// left, right := i, (i+1)%5

		// if i == 4 {
		// 	pH[i] = &Phil{
		// 		num:     i + 1,
		// 		leftCS:  cS[right],
		// 		rightCS: cS[left],
		// 	}
		// } else {

		// 	pH[i] = &Phil{
		// 		num:     i + 1,
		// 		leftCS:  cS[left],
		// 		rightCS: cS[right],
		// 	}
		// }

	}

	// //channel
	ch := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go pH[i].eat(ch, &wg)
	}

	// //done eating
	wg.Wait()
	close(ch)
}
func randomOrder(input []int) {
	fmt.Println("old", input)
	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
}
