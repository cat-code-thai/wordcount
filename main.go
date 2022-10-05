package main

import (
	"flag"
	"fmt"
	"github.com/cat-code-thai/wordcount/helper"
	"strings"
)

func main() {
	src := readInput()

	res := helper.CountWords(src)

	fmt.Println(res)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	flag.StringVar(&src, "p", "", "pattern to match against")
	flag.Parse()

	src = strings.Join(flag.Args(), "")

	return src
}
