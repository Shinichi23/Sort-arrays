package main

import (
	"fmt"
	"slices"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	merged := make([]int, 0)

	ints := make([]int, 12)
	fmt.Println("Enter a Int.")
	for i := 0; i < 12; i++ {
		fmt.Scan(&ints[i])
	}

	slice1 := ints[0:3]
	slice2 := ints[3:6]
	slice3 := ints[6:9]
	slice4 := ints[9:12]

	fmt.Println("Your subarray 1 before sorting: ", slice1)
	fmt.Println("Your subarray 2 before sorting: ", slice2)
	fmt.Println("Your subarray 3 before sorting: ", slice3)
	fmt.Println("Your subarray 4 before sorting: ", slice4)

	wg.Add(4)
	go sortArray(slice1, &wg)
	go sortArray(slice2, &wg)
	go sortArray(slice3, &wg)
	go sortArray(slice4, &wg)

	wg.Wait()

	fmt.Println("Your subarray sorted", slice1)
	fmt.Println("Your subarray sorted", slice2)
	fmt.Println("Your subarray sorted", slice3)
	fmt.Println("Your subarray sorted", slice4)

	merged = mergeArray(slice1, slice2, slice3, slice4)

	fmt.Println("Your array merged and sorted: ", merged)
}

func sortArray(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	slices.Sort(arr)
}

func mergeArray(arr1, arr2, arr3, arr4 []int) []int {
	merged := make([]int, 0)

	merged = append(merged, arr1...)
	merged = append(merged, arr2...)
	merged = append(merged, arr3...)
	merged = append(merged, arr4...)

	slices.Sort(merged)

	return merged
}
