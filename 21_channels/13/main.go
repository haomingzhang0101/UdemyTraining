package main

import (
	"fmt"
)

func gen(nums ...int) chan int  {
	out := make(chan int)
	go func() {
		for _, n := range nums{
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int  {
	out := make(chan int)
	go func() {
		for n := range in{
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main()  {
	in := gen(2, 3)
	//FAN OUT
	c1 := sq(in)
	c2 := sq(in)

	//FAN IN
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func merge(cs ...chan int) chan int  {
	out := make(chan int)
	//var wg sync.WaitGroup
	done := make(chan bool)
	//wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch{
				out <- n
			}
			//wg.Done()
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
