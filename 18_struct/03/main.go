package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main()  {
	p := person{"David", "Zhang", 22}
	fmt.Println(p.fullName())
}