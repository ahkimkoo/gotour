package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	m := float64(0)
	for c := 1; m-z < 0 || m-z > 0.0001; c++ {
		m = z
		z = z - (z*z-x)/2*z
		fmt.Printf("Round %v: %g\n", c, z)
	}
	return z
}
func main() {
	fmt.Println("result is ", Sqrt(2))
}
