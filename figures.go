package main

import (
	"math"
	"sync"
)

type point struct {
	x, y float64
}

func (p point) toScreen(centerX, centerY int, radius float64) (int, int) {
	return centerX + int(radius*p.x), centerY + int(radius*p.y)
}

// return x, y from -1 to 1
type figure func(time float64) []point

func circle(time float64) []point {
	x := math.Cos(time)
	y := math.Sin(time)
	return []point{{x, y}}
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

func (st *figureState) getCoords(time float64) []point {
	st.changeMut.Lock()
	points := st.coordsFunc(time)
	st.changeMut.Unlock()
	return points
}
