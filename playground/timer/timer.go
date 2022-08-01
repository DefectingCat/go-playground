package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	startTime time.Time
	endTime   time.Time
	diff      time.Duration
}

func (t *Timer) Start() {
	fmt.Println("Start record time.")
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	t.endTime = time.Now()
	t.diff = t.endTime.Sub(t.startTime)
	fmt.Println("Timer end, total time take: ", t.diff)
}
