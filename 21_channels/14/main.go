package main

import (
	"fmt"
)

func main() {

	in := gen()

	//FAN OUT
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)
	f10 := factorial(in)

	//FAN IN
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
		fmt.Println(n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			for j := 30; j < 40; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...chan int) chan int  {
	out := make(chan int)
	done := make(chan bool)

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch{
				out <- n
			}
			done <- true
		}(c)
	}

	go func() {
		for i := 0; i < len(cs); i++{
			<-done
		}
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern
*/