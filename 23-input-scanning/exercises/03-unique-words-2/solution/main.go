package main

// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	rx := regexp.MustCompile(`[^A-Za-z]+`)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++

		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)
		words[word]++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}
