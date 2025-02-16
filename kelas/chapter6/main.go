package main

import (
	"fmt"
	"math"
)

type CentiMeters float64

func (c CentiMeters) IsTooLong() {
	if c > 100 {
		fmt.Printf("besar!!")
	} else {
		fmt.Printf("pas")
	}
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	getArea() float64
	getParameter() float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getParameter() float64 {
	return 2*r.width + 2*r.height
}

func meansureShape(s shape) {
	fmt.Printf("shape has an area of %f\n", s.getArea())
	fmt.Printf("shape has an Parameter of %f\n", s.getParameter())
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getParameter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	width float64
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func (s square) getParameter() float64 {
	return 4 * s.width
}

func main() {
	myVar := CentiMeters(3.2)
	fmt.Printf("Type : %T , value  %v\n ", myVar, myVar)
	// myVar.IsTooLong()

	myRectangle := rectangle{
		width:  77,
		height: 32,
	}

	myCircle := circle{
		radius: 9,
	}

	mySquare := square{
		width: 8,
	}
	fmt.Printf("Type = %T, value %+v\n", myRectangle, myRectangle)
	meansureShape(myRectangle)
	meansureShape(myCircle)
	meansureShape(mySquare)

}
