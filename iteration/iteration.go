package iteration

import (
	"strings"
)

// RepeatCharacter loops through a character 5 times
func RepeatCharacter(s string, l int) string {

	cl := make([]string, l)
	for i := 0; i < len(cl); i++ {
		cl[i] = s
	}
	return strings.Join(cl, "")
}
