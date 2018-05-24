package main

import (
	"sort"
	"fmt"
)


func main()  {
	studentGroup := []string {"David", "Frank", "Sam", "Abby", "Bill"}

	sort.Slice(studentGroup, func(i, j int) bool {
		return studentGroup[i] > studentGroup[j]
	})

	fmt.Println(studentGroup)
}
