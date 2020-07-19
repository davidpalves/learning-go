package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper is a interface for the Sleep method
type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown is a function to count from the given number to 0
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

// ConfigurableSleeper is a struct to manage the Sleep() function settings
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep is a sleep implementation using ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
