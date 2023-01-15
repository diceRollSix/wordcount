package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("%#v", src)
	//fmt.Println()

	srcRunes := []rune(src)
	//fmt.Println(srcRunes)

	space := rune(' ')
	prevChar := space
	countWords := 0
	for _, val := range srcRunes {
		if val == space && prevChar != space {
			countWords++
		}
		prevChar = val
	}

	fmt.Println(countWords)
}

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}
