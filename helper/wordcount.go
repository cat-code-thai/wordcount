package helper

import (
	"strings"
)

func CountWords(str string) int {
	if str == "" {
		return 0
	}
	s := strings.Split(str, " ")

	return len(s)
}
