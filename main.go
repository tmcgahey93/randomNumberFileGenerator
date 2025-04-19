package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate_numbers.go <size>")
		return
	}

	// Parse size from command-line argument
	size, err := strconv.Atoi(os.Args[1])
	if err != nil || size <= 0 {
		fmt.Println("Please provide a valid positive integer for size.")
		return
	}

	// Generate numbers from 1 to size
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i + 1
	}

	// Shuffle the numbers
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// Create output file
	file, err := os.Create("numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write to file
	for _, num := range nums {
		fmt.Fprintln(file, num)
	}

	fmt.Printf("Generated numbers.txt with %d unordered numbers.\n", size)
}
