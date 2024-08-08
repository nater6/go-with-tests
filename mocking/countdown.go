package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type slySleeperOperations struct {
	Calls []string
}

const sleep = "sleep"

func (s *slySleeperOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

const write = "write"

func (s *slySleeperOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

var countdownStart = 3
var finalWord = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	sleeper := ConfigurableSleeper{duration: time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, &sleeper)
}
