package square
//package main

//import (
//	"fmt"
//)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	// implement me
	var e Point
	e.x = s.start.x + int(s.a)
	e.y = s.start.y + int(s.a)
	return e
}

func (s Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	// implement me
	return s.a * 4
}

//func main(){
//	var s Square = Square{start: Point{x: -2, y: 4}, a: 0}
//	e := s.End()
//	fmt.Println(e)
//	a := s.Area()
//	fmt.Println(a)
//	p := s.Perimeter()
//	fmt.Println(p)
//}
