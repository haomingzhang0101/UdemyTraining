package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	person
	first string
	LicenseToKill bool
}

func main()  {
	p := DoubleZero{
		person{
			"James",
			"Bond",
			20,
		},
		"Double zero seven",
		true,
	}

	fmt.Println(p.first, p.person.first)
}

