package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Define interface that 'real' sleeper and 'spy' sleeper satisfy
type Sleeper interface {
	Sleep()
}

// SpySleeper has a 'Calls' property
type SpySleeper struct {
	Calls int
}

// Empty type that will be used by main func
type DefaultSleeper struct {
}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	// Below has same signature as time.Sleep
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// Actual sleep function, which adds 1 sec sleep to main func
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// SpySleeper which simply counts numbers of calls, used by test
func (s *SpySleeper) Sleep() {
	s.Calls ++
}

// Countdown takes a Sleeper interface so we can invoke with
// either actual Default.Sleep() function or SpySleeper.Sleep()
func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}