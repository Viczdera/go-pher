package main

import (
	"fmt"
	"strconv"
	"strings"
)

func scanNumbers() ([]int, error) {
	var numbersInput string
	var numbersAsStrings []string
	fmt.Scan(&numbersInput)

	numbersAsStrings = strings.Split(numbersInput, ",")

	var numbersAsInt []int = make([]int, 10)
	for i := 0; i < 10 && i < len(numbersAsStrings); i++ {
		var number string = numbersAsStrings[i]
		intNumber, err := strconv.Atoi(number)

		if err != nil {
			return numbersAsInt, err
		}

		numbersAsInt[i] = intNumber
	}

	return numbersAsInt[:10], nil
}

func Swap(numbers []int, index int) {
	if index < len(numbers)-1 {
		temp := numbers[index]
		numbers[index] = numbers[index+1]
		numbers[index+1] = temp
	}
}

func BubbleSort(numbers []int) []int {
	var swapped bool = true
	for swapped == true {
		swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}
	}

	return numbers
}

func main() {
	fmt.Printf("Please enter up to ten numbers, comma-separated. Example:\n")
	fmt.Printf("1,2,3,4,5\n")
	fmt.Printf("Please note that the last number should not be followed by a comma ','\n")
	fmt.Printf("Enter your numbers:\n")

	scannedNumbers, err := scanNumbers()

	if err != nil {
		fmt.Printf("There was an error reading your numbers. Please check your input and re-try.\n")
	}

	BubbleSort(scannedNumbers)

	fmt.Printf("Your sorted input is: %d\n", scannedNumbers)
}
