package main

import (
	"sort"
	"fmt"
)

func main()  {
	studentGroup := []string {"David", "Frank", "Sam", "Abby", "Bill"}
	sort.StringSlice(studentGroup).Sort()
	//sort.Sort(sort.StringSlice(studentGroup))
	fmt.Println(studentGroup)
}
