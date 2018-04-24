package main

import "fmt"

func main()  {
	switch "Haoming"{
	case "Sam":
		fmt.Println("What's up Sam!")
	case "Haoming":
		fmt.Println("Nihao Haoming!")
	case "Michael", "Simmon":
		fmt.Println("Hallo Simmon or Michael!")
	default:
		fmt.Println("No friend")
	}
}
