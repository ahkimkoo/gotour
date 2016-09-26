package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //指针不可以直接赋值，但是结构体指针字段可以直接修改
	fmt.Println(v)
	*p = Vertex{3, 4}
	fmt.Println(v)
}
