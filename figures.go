package main

import (
	"math"
	"sync"
)

// return x, y from -1 to 1
type figure func(time float64) (float64, float64)

func circle(time float64) (float64, float64) {
	x := math.Cos(time)
	y := math.Sin(time)
	return x, y
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

func (st *figureState) getCoords(time float64) (float64, float64) {
	st.changeMut.Lock()
	x, y := st.coordsFunc(time)
	st.changeMut.Unlock()
	return x, y
}
