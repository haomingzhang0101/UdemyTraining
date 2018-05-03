package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	p := &person{"David", 22}
	fmt.Println(p.name)
	fmt.Println(p.age)
}
