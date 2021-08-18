package utils

import "fmt"

func init() {
	fmt.Println("utils init!")
}

var (
	a = 1
	b = 2
	C = 42
)

func SayHello(name string) {
	fmt.Printf("Hello, %v!", name)
}

func Fibonacci(n int, a int, b int) int {
	if n <= 0 {
		return b
	}

	return Fibonacci(n-1, b, a+b)
}
