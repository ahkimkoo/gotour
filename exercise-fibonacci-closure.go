package main

import "fmt"

func fibonacci() func() int {
	p2, p1 := 0, 1
	return func() int {
		sum := p1 + p2
		p2 = p1
		p1 = sum
		return p1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
