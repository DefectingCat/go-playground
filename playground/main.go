package main

import (
	"fmt"
	"math"
	"time"
	"timer/timer"
)

func add(x, y float64) float64 {
	return x + y
}

func swap(a, b int) (int, int) {
	return b, a
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(oldZ-z) <= 0.00000001 {
			break
		}
	}
	return z
}

func fibonacci22(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func fibonacci2(n int) int {
	x, y := 0, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

func fibo2(n int) int {
	if n < 2 {
		return n
	}
	return fibo2(n-1) + fibo2(n-2)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v", len(x), cap(x), x)
}

type Phone interface {
	call()
}
type iPhone struct {
	name string
}

func (iPhone *iPhone) call() string {
	fmt.Println("phone")
	return iPhone.name
}

func returnNil() *int {
	n := 3
	return &n
}

func findNum(num int, nums ...int) {
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println("fount at index: ", i)
			found = true
		}
	}
	if !found {
		fmt.Println("not found")
	}
}

func change(s *[]string) {
	(*s)[0] = "xfy"
	*s = append(*s, "dfy")
	fmt.Println(*s)
}

func showNums() {
	for i := range make([]int, 10) {
		time.Sleep(10)
		fmt.Print(i)
	}
}
func showChar() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(10)
		fmt.Printf("%c", i)
	}
}

func hello(done chan bool) {
	fmt.Println("Hello go routine")
	done <- true
}

func channelSqrt(x float64, c chan float64) {
	z := 1.0
	for {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(oldZ-z) <= 0.00000001 {
			break
		}
	}
	c <- z
}

func useChannel() {
	var t timer.Timer
	t.Start()

	c := make(chan float64)

	times := 100
	for i := 1; i <= times; i++ {
		go channelSqrt(float64(i), c)
	}

	var results []float64
	for i := 1; i <= times; i++ {
		results = append(results, <-c)
	}
	fmt.Println(results)
	t.Stop()
}
func channelTest() {
	var t timer.Timer
	t.Start()

	times := 100
	var results []float64
	for i := 1; i <= times; i++ {
		results = append(results, Sqrt(float64(i)))
	}
	fmt.Println(results)

	t.Stop()
}

func main() {
	channelTest()
}
