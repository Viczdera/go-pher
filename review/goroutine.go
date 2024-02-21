package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	numThreads = 4
)

// This is the goroutine that will be called multiple times to sort a slice. The slice is copied at
// first in order to leave the original slice unchanged. Then, the copy of the slice is sorted in
// place and the sorted slice is returned via a channel.
func sortIt(unsorted []int, idx int, c chan []int) {
	fmt.Printf("Unsorted slice, goroutine %d: %v\n", idx, unsorted)
	result := make([]int, len(unsorted))
	copy(result, unsorted)
	sort.Ints(result)
	c <- result
}

// Convert a slice of strings to a slice of ints until the first conversion error occurs. The
// integer slice and the error (can be nil) are returned.
func mapStringsToInts(stringSlice []string) ([]int, error) {
	result := make([]int, 0, len(stringSlice))
	var err error
	// Actually convert each field to an integer and append each to the slice.
	// Do not continue after the first conversion error.
	for count := 0; count < len(stringSlice) && err == nil; count++ {
		var input int
		input, err = strconv.Atoi(stringSlice[count])
		if err == nil {
			result = append(result, input)
		}
	}
	return result, err
}

// Read a single line from stdin. Note that defining `scanner` in this function causes any input to
// stdin after the first line to be lost to this program. Since we only ever want to read in a
// single line, that is not a problem in this case.
func readLineFromStdin() (string, error) {
	var scanner = bufio.NewScanner(os.Stdin)
	// The variable `err` will hold the error determined by `scanner`, if any.
	var err error
	// Advance the scanner to the next newline character.
	if !scanner.Scan() {
		err = scanner.Err()
	}
	// Put whatever could be read in into `line`.
	line := scanner.Text()
	return line, err
}

// Read a slice of integers from a single line from standard input.
func getUserInput() ([]int, error) {
	line, err := readLineFromStdin()
	fields := strings.Fields(line)
	// Convert the strings to ints, but only if there hasn't been an error yet.
	var slice []int
	if err == nil {
		slice, err = mapStringsToInts(fields)
	}
	return slice, err
}

// Compute the sum over all integers in a slice.
func sum(input []int) int {
	result := 0
	for _, val := range input {
		result += val
	}
	return result
}

// Partition a slice into a number of shorter slices, all of which have about the same length. The
// resulting slices have exactly the same length if the length of the original slice is divisible by
// the number of partitions. Otherwise some of the first slices are one element longer.
func partition(partitions int, sli []int) [][]int {
	// Compute the expected size for each of the partitions.
	// First, get the average size.
	avgSize := int(len(sli) / partitions)
	// Then, determine whether any of the partitions need to be one element longer to be able to
	// contain all elements.
	elems := make([]int, partitions)
	for idx := range elems {
		elems[idx] = avgSize
	}
	// The first slices will always be the longer ones.
	mismatch := len(sli) - avgSize*partitions
	for idx := 0; idx < mismatch; idx++ {
		elems[idx]++
	}
	// Treat cases where we would have empty partitions. Throw away those partitions that are not
	// needed. Technically, this causes us to use fewer than the requested 4 goroutines to sort
	// slices with fewer than 4 elements, but I don't think that is a problem.
	for idx, val := range elems {
		if val == 0 {
			partitions = idx
			break
		}
	}
	elems = elems[:partitions]
	// Actually partition the original slice into sub-slices.
	result := make([][]int, partitions)
	for idx := 0; idx < partitions; idx++ {
		startIdx := sum(elems[:idx])
		endIdx := sum(elems[:idx+1])
		// This uses a way to access sub-slices that sets both length and capacity. That is,
		// accessing a slice of c elements like `slice[a:b]` will result in a slice with `b-a`
		// elements and a capacity of `c-b`. That means appending to the resulting slice can
		// potentially overwrite elements. Accessing it like `slice[a:b:b]` will set the capacity to
		// the length and, thus, avoid problems due to overwriting.
		result[idx] = sli[startIdx:endIdx:endIdx]
	}
	return result
}

// Determines the maximum difference between elements in the slices associated with the same index.
// This is not the maximum of the absolute values but the maximum of the actual differences.  That
// is, it computes the difference between the i-th element of `sli1` and the i-th element of `sli2`
// and returns the largest one.
// Ignores length differences of the slices, i.e. handles them the same way as Python's zip.
func maxIntSliceDiff(sli1, sli2 []int) int {
	// Using a starting value of 0 means that we return a difference of 0 in case one slice is
	// empty.
	result := 0
	// Determine minimum length of the available slices.
	minLen := len(sli1)
	if len(sli2) < minLen {
		minLen = len(sli2)
	}
	// Initialize the starting value.
	if minLen > 0 {
		result = sli1[0] - sli2[0]
	}
	// Determine the maximum difference.
	for idx := 0; idx < minLen; idx++ {
		diff := sli1[idx] - sli2[idx]
		if diff > result {
			result = diff
		}
	}
	return result
}

// This determines the smallest element from the slices in `values` whose places in their respective
// slices are given by `indices`. Indices too large for a slice are ignored.
// Ignores length differences of the slices, i.e. handles them the same way as Python's zip.
func smallestElement(values [][]int, indices []int) (element, index int) {
	// An index of -1 indicates no element could be found, for example because all values in
	// `indices` are out of bounds for the slices in `values`.
	index = -1
	// Determine minimum length of the available slices.
	minLen := len(values)
	if len(indices) < minLen {
		minLen = len(indices)
	}
	for idx := 0; idx < minLen; idx++ {
		// Ignore indices that are out of bounds for the associated slice.
		if indices[idx] >= len(values[idx]) {
			continue
		}
		if index == -1 || values[idx][indices[idx]] < element {
			element = values[idx][indices[idx]]
			index = idx
		}
	}
	return
}

// Merge a slice of integer slices in ascending order. This assumes each of the input slices is
// sorted in ascending order.
func mergeAscending(input [][]int) []int {
	// The variable `counter` will move from 0 to the length of the respective sorted slice.
	counter := make([]int, len(input))
	// The variable `maxCounter` will contain the stop value for each value in `counter`.
	maxCounter := make([]int, len(input))
	for idx, sli := range input {
		maxCounter[idx] = len(sli)
	}
	result := []int{}
	// We are not done until each of the indices has reached its maximum value.
	for maxIntSliceDiff(maxCounter, counter) != 0 {
		// Pick the smallest element from all the slices in `input` that the values in `counter`
		// point to.
		newVal, incIdx := smallestElement(input, counter)
		if incIdx >= 0 {
			result = append(result, newVal)
			counter[incIdx]++
		}
	}
	return result
}

// Implement a limited form of merge sort. In merge sort, the list is successively split up into
// smaller and smaller chunks that are then merged in sorted order. Here, we only split once into
// partitions, sort each of the partitions, and then merge them in ascending order. The term
// "limited" is chosen since we only split once.
// See here https://en.wikipedia.org/wiki/Merge_sort for a description of vanilla merge sort.
func limitedMergeSort(input []int) []int {
	partitioned := partition(numThreads, input)
	c := make(chan []int, numThreads)
	for idx, subslice := range partitioned {
		go sortIt(subslice, idx+1, c)
	}
	result := make([][]int, len(partitioned))
	for idx := 0; idx < len(partitioned); idx++ {
		result[idx] = <-c
	}
	return mergeAscending(result)
}

func main() {
	var input []int
	for done := false; !done; {
		var err error
		input, err = getUserInput()
		done = err == nil
	}
	// The task did not ask for printing the entire unsorted slice.
	// fmt.Printf("Unsorted slice: %v\n", input)
	result := limitedMergeSort(input)
	fmt.Printf("Sorted slice: %v\n", result)
}
