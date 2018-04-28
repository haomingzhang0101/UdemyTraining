package main

import "fmt"

func main()  {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "早上好"
	greeting = append(greeting, "Guten Morgan")

	fmt.Println(greeting[3])
}
