package helper

import (
	"strconv"
	"strings"
)

func CountWords(str string) string {
	if str == "" {
		return "0"
	}
	s := strings.Split(str, " ")

	return strconv.Itoa(len(s))
}
