package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func incrementor(s string)  {
	for i := 0; i < 20; i++  {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		//x := counter
		//x++
		//counter = x
		counter++
		fmt.Println(s, i, "Counter", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func main()  {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}