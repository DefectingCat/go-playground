package channel_test

import (
	"fmt"
	"math"
	"timer/timer"
)

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

func channel_test() {
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
	channel_test()
}
