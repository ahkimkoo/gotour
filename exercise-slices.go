package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slices := make([][]uint8, dy)
	for x := range slices {
		for y := 0; y < dx; y++ {
			slices[x] = append(slices[x], uint8(x^y))
		}
	}
	return slices
}

func main() {
	pic.Show(Pic)
}
