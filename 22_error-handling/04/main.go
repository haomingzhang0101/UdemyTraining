package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, lon string
	err error
}

func (n *NorgateMathError) Error() string  {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.lon, n.err)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("error: square root of a negative number:%v", f)
		return 0, &NorgateMathError{"50 N", "99 W", nme}
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