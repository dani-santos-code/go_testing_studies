package main

import (
	"os"

	"github.com/dani-santos-code/go_testing_studies/countdown"
)

func main() {
	countdown.Countdown(os.Stdout, &countdown.DefaultSleeper{})
}
