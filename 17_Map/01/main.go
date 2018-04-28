package main

import "fmt"

func main()  {
	greeting := make(map[string]string)
	greeting["David"] = "Good morning"
	greeting["Jean"] = "早上好"

	fmt.Println(greeting)
}
