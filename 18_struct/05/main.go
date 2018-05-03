package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last string `json:"-"`
	Age int `json:"myAge"`
	notExported int
}

func main()  {
	p1 := Person{"Haoming", "Zhang", 22, 5955}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}