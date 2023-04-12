package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
