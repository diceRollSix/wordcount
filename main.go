package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	src := readInput()
	//fmt.Printf("%#v", src)
	//fmt.Println()

	srcRunes := []rune(src)
	//fmt.Println(srcRunes)

	space := ' '
	prevChar := space
	countWords := 0
	for _, val := range srcRunes {
		if val == space && prevChar != space {
			countWords++
		}
		prevChar = val
	}

	if prevChar != space {
		countWords++
	}

	fmt.Println(countWords)
}

func readInput() string {
	flag.Parse()
	return strings.Join(flag.Args(), "")
}
