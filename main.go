package main

import (
	"fmt"
)

func main() {
	actualSize := 6
	arr := []int{1, 2, 3, 4, 5}

	missingNum := FindMissing(arr, actualSize)
	fmt.Printf("%d", missingNum)
}

//FindMissing finds the missing number
func FindMissing(arr []int, actualSize int) int {
	//Use AP to find sum of digits
	sum := FindSum(actualSize)
	for i := 0; i < len(arr); i++ {
		sum = sum - arr[i]
	}
	return sum
}

//FindSum Uses AP by formulat n/2(1+n)
func FindSum(size int) int {
	return (size * (1 + size)) / 2
}
