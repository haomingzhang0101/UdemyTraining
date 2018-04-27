package main

import "fmt"

func max(sf ...int) int {
	var max int
	for i, v := range sf{
		if v > max || i == 0 {
			max = v
		}
	}
	return max
}

func main()  {
	fmt.Println(max(-1, -2, -3))
}

