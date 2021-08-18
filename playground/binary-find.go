package main

import (
	"fmt"
	"math"
)

func binaryFind(arr *[]int, leftIndex int, rightIndex int, target int) int {
	if leftIndex > rightIndex {
		fmt.Println("Not Found!")
		return -1
	}

	midIndex := int(math.Round(float64((leftIndex + rightIndex) / 2)))

	if (*arr)[midIndex] > target {
		midIndex = binaryFind(arr, leftIndex, midIndex-1, target)
	} else if (*arr)[midIndex] < target {
		midIndex = binaryFind(arr, midIndex+1, rightIndex, target)
	} else {
		fmt.Println("Target index is", midIndex)
		return midIndex
	}
	return midIndex
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := binaryFind(&arr, 0, len(arr)-1, 6)
	fmt.Println(arr[i])
}
