package main

import (
	"math"
	"sync"
)

type floatPoint struct {
	x, y float64
}

type intPoint struct {
	x, y int
}

func (p floatPoint) toScreen(center intPoint, radius float64) intPoint {
	return intPoint{center.x + int(radius*p.x), center.y + int(radius*p.y)}
}

// return x, y from -1 to 1
type figure func(time float64) []floatPoint

func circle(time float64) []floatPoint {
	x := math.Cos(time)
	y := math.Sin(time)
	return []floatPoint{{x, y}}
}

type figureState struct {
	coordsFunc figure
	changeMut  sync.Mutex
}

func (st *figureState) change(newFigure figure) {
	st.changeMut.Lock()
	st.coordsFunc = newFigure
	st.changeMut.Unlock()
}

func (st *figureState) getCoords(time float64) []floatPoint {
	st.changeMut.Lock()
	points := st.coordsFunc(time)
	st.changeMut.Unlock()
	return points
}
