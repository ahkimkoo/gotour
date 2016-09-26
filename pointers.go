package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i //var p *int =
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	fmt.Println(p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
