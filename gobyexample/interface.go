package main

import (
	"fmt"
	"math"
)

// 几何接口
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 必须实现接口的所有方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
