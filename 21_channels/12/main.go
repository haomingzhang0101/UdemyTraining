package main

import "fmt"

func factorial(in chan int) chan int  {
	out := make(chan int)
	go func() {
		for n := range in{
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func gen() chan int  {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10 ; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fact(n int) int  {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func main()  {
	in := gen()
	out := factorial(in)
	for n := range out {
		fmt.Println(n)
	}
}