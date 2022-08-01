package basic

import (
	"fmt"
	"math"
)

func Struct_() {
	//구조체
	//dic := newDict()
	//dic.data[1] = "A"
	//println(dic.data)

	//메서드
	//rect := Rect{10, 20}
	//println(rect.area())
	//println(rect.area2())
	//println(rect.area()) //area2로 인해 값이 변경

	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"
	printIt(x)

	var a interface{} = 100
	i := a
	j := a.(int) // type assertion

	println(i)
	println(j)
}

//구조체
type dict struct {
	data map[int]string
}

func newDict() *dict { //포인터 반환
	d := dict{}
	d.data = map[int]string{}
	return &d //주소값 반환 = 포인터 위치 반환
}

//메서드
type Rect struct {
	width, height float64
}

func (r *Rect) area2() float64 {
	r.width++
	return r.width * r.height
}

//interface

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shape ...Shape) {
	for _, s := range shape {
		a := s.area()
		println(a)
	}
}

//empty interface
func printIt(v interface{}) {
	fmt.Println(v)
}
