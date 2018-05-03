package main

import "fmt"

func main()  {
	var stu []string
	stu = append(stu, "hm")
	fmt.Println(len(stu))
	fmt.Println(cap(stu))
	stu = append(stu, "hm1")
	fmt.Println(len(stu))
	fmt.Println(cap(stu))
	stu = append(stu, "hm2")
	fmt.Println(len(stu))
	fmt.Println(cap(stu))
}
