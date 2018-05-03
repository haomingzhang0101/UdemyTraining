package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int `json:"myAge"`
}

func main()  {
	var p Person
	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)

	bs := []byte(`{"First":"Haoming", "Last":"Zhang", "myAge":22}`)
	json.Unmarshal(bs, &p)
	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)
}