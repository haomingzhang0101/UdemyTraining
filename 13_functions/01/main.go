package main

import "fmt"

func main()  {
	fmt.Println(greet("David", "Zhang"))
}

func greet(fname, lname string) string  {
	return fmt.Sprint(fname," ", lname)
}

func greet2(fname, lname string) (string, string)  {
	return fmt.Sprint(fname," ", lname), fmt.Sprint(lname," ", fname)
}
