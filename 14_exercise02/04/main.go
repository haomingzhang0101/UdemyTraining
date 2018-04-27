package main

import "fmt"

func foo(sf ...int)  {
	fmt.Println(sf)
}

func main()  {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
