// Name: Sean Duncan
// NetID: sdduncan
// Description: The solution to Question 2 of Assignment 1-1. This file sums the numbers in a
// file concurrently
package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	var sum int
	for num := range nums {
		sum += num
	}
	out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	file, err := os.Open(fileName)
	defer file.Close()
	checkError(err)

	values, err := readInts(file)
	checkError(err)

	numChan := make(chan int, len(values))
	outChan := make(chan int, num)
	defer close(outChan)


	// sync up goroutines via wait group. You are


	// Spawn num goroutines and use a wait group to ensure
	// all goroutines finish before taking the
	for i := 0; i < num; i++ {
		go sumWorker(numChan, outChan)
	}

	// Send all values to numChan and close when done
	for _, value := range values {
		numChan <- value
	}
	close(numChan)

	var sum int
	for i := 0; i < num; i++{
		sum += <-outChan
	}

	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
