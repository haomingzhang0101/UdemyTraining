package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//b := &a
	var b *int = &a //referencing (point to a memory address where an int is stored)

	fmt.Println(b)
	fmt.Println(*b) //dereferencing

	*b = 42
	fmt.Println(a)
}

