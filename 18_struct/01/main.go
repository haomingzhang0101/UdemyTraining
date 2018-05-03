package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main()  {
	p := person{"Haoming", "Zhang", 22}
	fmt.Println(p.first, p.last, p.age)
}
