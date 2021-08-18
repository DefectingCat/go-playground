package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	arr := []int{1, 2, 34, 7, 5, 34, 54, 5}

	c := make(chan int)
	go sum(arr[:len(arr)/2], c)
	go sum(arr[len(arr)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
