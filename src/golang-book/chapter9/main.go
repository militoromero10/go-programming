package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{x: 0, y: 0, r: 1.25}
	r := Rectangle{x1: 0, x2: 3, y1: 0, y2: 3}
	fmt.Println("Area del circulo", c, c.area())
	fmt.Println("Area del rectangulo", r, r.area())
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	h := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (h + w)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
	perimeter() float64
}
