package countdown

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper interface
type Sleeper interface {
	Sleep()
}

//SpySleeper struct
type SpySleeper struct {
	Calls int
}

// Sleep function
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper struct
type DefaultSleeper struct{}

// Sleep function
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Countdown prints to stdout
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
