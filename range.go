package main

import "fmt"

var pow = []int{1, 2, 4, 8, 18, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
		x := pow[i : i+1]
		x[0] *= 3
	}
	fmt.Println(pow)
}
