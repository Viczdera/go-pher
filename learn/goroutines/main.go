package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("goroutines")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s)
	}
}
