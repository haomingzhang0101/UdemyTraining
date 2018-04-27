package main

import "fmt"

func changeMe(z map[string]int)  {
	z["Haoming"] = 22
}

func main()  {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Haoming"])
}
