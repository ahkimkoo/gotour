package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "你好"
	a[1] = "世界"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
