package main

import (
	"math"
	"fmt"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s *circle) area() float64  {
	return math.Pi * s.radius * s.radius
}

func info(s shape)  {
	fmt.Println("area is:", s.area())
}

func main()  {
	c := circle{3}
	fmt.Println(&c)
	info(&c)
}
