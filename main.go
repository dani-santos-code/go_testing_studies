package main

import "github.com/dani-santos-code/go_testing_studies/countdown"
import "os"


func main() {
	countdown.Countdown(os.Stdout, &countdown.DefaultSleeper{})
}
