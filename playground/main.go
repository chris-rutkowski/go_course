package main

import "fmt"

type shape interface {
    getArea() float64
}

type triangle struct {
    height float64
    base   float64
}

type square struct {
    sideLength float64
}

func main() {
    t := triangle{base: 10, height: 20}
    s := square{sideLength: 15}

    printArea(t)
    printArea(s)
}

func printArea(s shape) {
    fmt.Printf("Area is: %v\n", s.getArea())
}

func (t triangle) getArea() float64 {
    return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
    return s.sideLength * s.sideLength
}