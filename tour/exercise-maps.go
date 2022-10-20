package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	sentence := strings.Fields(s)
	for _, word := range sentence {
		ans[word]++
	}
	return ans
}

func main() {
	wc.Test(WordCount)
}
