package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		_, es := m[w]
		if es {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
