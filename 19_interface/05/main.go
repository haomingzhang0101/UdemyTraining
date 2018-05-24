package main

import (
	"sort"
	"fmt"
)

func main()  {
	n := []int {5, 3, 1, 8, 2}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
