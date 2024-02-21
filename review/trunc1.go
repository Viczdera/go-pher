package main

import (
	"fmt"
)

func main() {
	var number float64

	fmt.Println("Enter a float:")
	fmt.Scanf("%f", &number)

	var integer int = int(number)

	fmt.Println(integer)
}
