package main

import "fmt"

const PI = 3.14

type Shape struct {
	x, y float64
}

type Rectangle struct {
	Shape
	w, h float64
}

type Square struct {
	Rectangle
	b float64
}

type Circle struct {
	Shape
	r float64
}

func main() {
	var re Rectangle
	fmt.Scan(&re.w, &re.h)
	fmt.Println(re.GetArea())
	var c Circle
	fmt.Scan(&c.r)
	fmt.Println(c.GetArea())
	var s Square
	fmt.Scan(&s.b)
	fmt.Println(s.GetArea())
}

func (re Rectangle) GetArea() float64 {
	return re.w * re.h
}

func (c Circle) GetArea() float64 {
	return c.r * c.r * PI
}

func (s Square) GetArea() float64 {
	return s.b * s.b
}

/*
 * 运行时间：4ms 超过12.50%用Go提交的代码
 * 占用内存：960KB 超过12.50%用Go提交的代码
**/
