package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
				continue
			}
		}
		close(out)
	}()
	return out
}

func main() {

	c1 := gen(1, 2, 3)
	c2 := gen(1, 2, 3)
	out := make(chan int)

	Merge2Channels(ff, c1, c2, out, 3)

	for i := 0; i < 3; i++ {
		select {
		case n := <-out:
			fmt.Println(n)
		}
	}
}

func ff(a int) int {
	return a * a
}

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	for i := 0; i < n; i++ {
		out1 := exec(f, in1)
		out2 := exec(f, in2)
		sum(out1, out2, out)
	}
}

func exec(f func(int) int, in1 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		n0 := <-in1
		out <- f(n0)
		close(out)
	}()
	return out
}

func sum(in1 <-chan int, in2 <-chan int, out chan<- int) {
	go func() {
		n1 := <-in1
		n2 := <-in2
		out <- n1 + n2
	}()
}
