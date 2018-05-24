package main

import (
	"errors"
	"log"
)

var ErrSqrtError = errors.New("error: square root of a negative number")

func Sqrt(num float64) (float64, error)  {
	if num < 0 {
		return 0, ErrSqrtError
	}
	//To be implemented
	return 42, nil
}

func main()  {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
