package helper

import (
	"strings"
)

func CountWords(str string) int {
	s := strings.Split(str, " ")

	return len(s)
}
