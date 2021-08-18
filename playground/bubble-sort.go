package main

import "fmt"

func bubbleSort(arr *[]int) {
	n := len(*arr) - 1
	for i := n; i > 1; i-- {
		for j := 0; j < i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	arr := []int{12, 2, 13, 44, 1, 7, 14}

	bubbleSort(&arr)
	fmt.Println(arr)
}
