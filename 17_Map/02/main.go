package main

import (
	"fmt"
)

func main()  {
	
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!",
		2: "早上好",
	}
	
	fmt.Println(myGreeting)

	if val, ok := myGreeting[2]; ok {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", ok)
	}else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", ok)
	}
}
