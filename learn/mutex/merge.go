package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	fmt.Println("Enter array of integers e.g '2 32 3 1'")
	fmt.Println(">")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		fmt.Println(err)
		return
	}
	command = strings.TrimSpace(command)
	inputArr := strings.Split(command, " ")

	intArr, intErr := toIntArray(inputArr)
	if intErr != nil {
		fmt.Println("Use a single withspace only followed by numbers")
		return
	}

	stepSize := len(inputArr) / 4
	splitArr := splitArray(intArr, stepSize)

	//def wait group
	waitG := &sync.WaitGroup{}
	mutX := &sync.Mutex{}
	rSize := 4

	waitG.Add(4)

	for i := 0; i < rSize; i++ {
		go func(wg *sync.WaitGroup, mX *sync.Mutex, idX int) {
			mX.Lock()
			fmt.Printf("T%d: %v\n", idX+1, splitArr[idX])
			splitArr[idX] = bubbleSort(splitArr[idX])
			mX.Unlock()
			waitG.Done()
		}(waitG, mutX, i)
	}
	//notifying main that there's a wait group
	waitG.Wait()
	//main go routine
	mergedArr := merge(splitArr)
	fmt.Println(mergedArr)

}

func merge(arr [][]int) []int {
	merged := make([]int, 0, len(arr[0])+len(arr[1])+len(arr[2])+len(arr[3]))
	merged = append(merged, arr[0]...)
	merged = append(merged, arr[1]...)
	merged = append(merged, arr[2]...)
	merged = append(merged, arr[3]...)
	bubbleSort(merged)
	return merged
}
func splitArray(inputArr []int, stepSize int) [][]int {
	arr := [][]int{}
	for i := 0; i < 4; i++ { //since 1/4

		start := i * stepSize
		end := (i + 1) * stepSize
		//final division
		if i == 3 {
			arr = append(arr, inputArr[start:])
		} else {
			arr = append(arr, inputArr[start:end])
		}

	}
	return arr
}

func toIntArray(input []string) ([]int, error) {
	var intArr []int
	for _, strValue := range input {
		intValue, err := strconv.Atoi(strings.TrimSpace(strValue))
		if err != nil {
			return nil, err
		}
		intArr = append(intArr, intValue)
	}
	return intArr, nil
}

//sort nd swap

func bubbleSort(items []int) []int {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}
