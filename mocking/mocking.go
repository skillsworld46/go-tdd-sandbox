package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepFn  func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFn(s.Duration)
}

// type SpyTime struct {
// 	durationSlept time.Duration
// }

// func (s *SpyTime) Sleep(duration time.Duration) {
// 	s.durationSlept = duration
// }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(out, "Go!")
}
