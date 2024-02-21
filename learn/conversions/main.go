package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your bugger rating")
	fmt.Println("Rating should be in the range of 0 to 5")

	//comma ,ok syntax || comma,err syntax

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("Got it! Thanks for rating", input)
	//converting
	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Increased rating by 1. New rating is", numRating+1)
	}

}
