package main

import (
	fmt "fmt"
	math "math"
)

type Shape interface {
	Area() float64
}
type Rect struct {
	x float64
	y float64
	w float64
	h float64
}

func (p *Rect) Area() float64 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:12
	return p.w * p.h
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (p *Circle) Area() float64 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:20
	return math.Pi * p.r * p.r
}
func Area(shapes ...Shape) float64 {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:24
	s := 0.0
	for
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:25
	_, shape := range shapes {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:26
		s += shape.Area()
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:28
	return s
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:31
	rect := &Rect{0, 0, 2, 5}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:32
	circle := &Circle{0, 0, 3}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/33-Interface/shape.gop:33
	fmt.Println("area:", Area(circle, rect))
}
