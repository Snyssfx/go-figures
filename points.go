package main

type floatPoint struct {
	x, y float64
}

type intPoint struct {
	x, y int
}

func (p floatPoint) toScreen(center intPoint, radius float64) intPoint {
	return intPoint{center.x + int(radius*p.x), center.y + int(radius*p.y)}
}