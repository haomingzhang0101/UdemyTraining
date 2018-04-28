package main

import "fmt"

func main()  {
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Todd"
	student1[1] = "Mcleod"
	student1[2] = "100.00"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Stephen"
	student2[1] = "Grider"
	student2[2] = "100.00"

	records = append(records, student2)

	fmt.Println(records)
}
