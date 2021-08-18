package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// 循环发送数据给 channel
		c <- y
		x, y = y, x+y
	}
	// 显示关闭 channel
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	// 不断的读取 channel 里的数据，知道 channel 被关闭
	for i := range c {
		fmt.Println(i)
	}
}
