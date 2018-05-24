package main

import (
	"fmt"
	"log"
)

func Sqrt(num float64) (float64, error)  {
	if num < 0 {
		return 0, fmt.Errorf("error: square root of a negative number:%v", num)
	}
	//To be implemented
	return 42, nil
}


func main()  {
	_, err := Sqrt(-10.22)
	if err != nil {
		log.Println(err)
	}
}
