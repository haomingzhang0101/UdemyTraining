package main

import "fmt"

func main()  {

	mySlice := []string {"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("你好"[2])
}
