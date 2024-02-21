package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your bugger rating")

	//comma ,ok syntax || comma,err syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Got it!", input)

	fmt.Printf("Type of rating is %T", input)

}
