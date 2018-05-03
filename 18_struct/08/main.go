package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string
	Age int `json:"myAge"`
}

func main()  {
	var p Person
	rdr := strings.NewReader(`{"First":"Haoming", "Last":"Zhang", "myAge":22}`)
	json.NewDecoder(rdr).Decode(&p)

	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)
}
