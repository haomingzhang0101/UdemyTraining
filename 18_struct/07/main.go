package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last string `json:"-"`
	Age int `json:"myAge"`
	notExported int
}

func main()  {
	p := Person{"Haoming", "Zhang", 22, 5955}
	json.NewEncoder(os.Stdout).Encode(p)
}