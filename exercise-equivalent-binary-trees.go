package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println(t1)
	fmt.Println(t2)
	go func() {
		Walk(t1, ch1)
		ch1 <- 0
	}()

	go func() {
		Walk(t2, ch2)
		ch2 <- 0
	}()

	for {
		t1 := <-ch1
		t2 := <-ch2
		if t1 == 0 && t2 == 0 {
			return true
		}

		if t1 == t2 {
			continue
		} else {
			return false
		}
	}
}

func main() {
	c := make(chan int)
	go func() {
		Walk(tree.New(1), c)
		c <- 0
	}()
	for {
		t := <-c
		if t == 0 {
			break
		}
		fmt.Println(t)
	}
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
