package main

import "fmt"

func main() {
	mynum := 43

	var ptr = &mynum
	fmt.Println("value is", ptr)
}
