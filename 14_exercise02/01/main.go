package main

import "fmt"

func half(n int) (int, bool)  {
	return n / 2, n % 2 == 0
}

func main()  {
	n, res := half(2)
	fmt.Println(n, res)
}