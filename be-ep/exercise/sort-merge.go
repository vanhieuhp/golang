package main

import (
	"fmt"
	"sort"
	"sync"
)

// mergeSortedArrays merges multiple sorted slices into one sorted slice
func mergeSortedArrays(slices ...[]int) []int {
	result := []int{}
	for _, slice := range slices {
		result = append(result, slice...)
	}
	sort.Ints(result) // Final merge and sort
	return result
}

func main() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter", n, "integers:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Determine the size of each partition
	chunkSize := (n + 3) / 4 // To ensure we divide properly

	// Create slices for partitions
	partitions := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > n {
			end = n
		}
		partitions[i] = arr[start:end]
	}

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			sort.Ints(partitions[index]) // Sort the partition
			fmt.Printf("Sorted partition %d: %v\n", index+1, partitions[index])
		}(i)
	}

	wg.Wait() // Wait for all sorting to complete

	// Merge sorted partitions
	finalSortedArray := mergeSortedArrays(partitions[0], partitions[1], partitions[2], partitions[3])
	fmt.Println("Final sorted array:", finalSortedArray)
}
