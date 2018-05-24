package main

import (
	"os"
	"log"
	"fmt"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("Error handled: ", err)
		//log.Println("Error handled: ", err)
		//log.Fatalln(err)
		panic(err)
	}
}
