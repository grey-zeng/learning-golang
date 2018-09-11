package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	// 常规使用方式
	// point1 := Point{0, 0}
	// point2 := Point{1, 1}
	// fmt.Println(point1.Distance(point2))

	// 嵌套使用
	point1 := ColoredPoint{&Point{0, 0}, color.RGBA{255, 0, 0, 255}}
	point2 := ColoredPoint{&Point{1, 1}, color.RGBA{0, 0, 255, 255}}
	// 这里就可以这么使用
	fmt.Println(point2.X)
	// 必须使用强制的类型
	fmt.Println(point1.Distance(*point2.Point))

	// 方法值和方法表达式
	var distance = point1.Distance
	fmt.Println(distance(*point2.Point))
	var tDistance = ColoredPoint.Distance
	fmt.Println(tDistance(point1, *point2.Point))
}
