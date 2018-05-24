package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo()  {
	for i := 0; i < 45; i++  {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 45; i++  {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func main()  {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}