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

func sinWithX(time float64) []floatPoint {
	y := math.Sin(1.5 * time)

	twoPiInX := math.Floor(time / (2 * math.Pi))
	x := (time-twoPiInX*2*math.Pi)/(1*math.Pi) - 1

	return []floatPoint{
		{x, y},
		{x, 0},
	}
}

func itsAllCos(time float64) []floatPoint {
	cos := math.Cos(1.0 * time)

	twoPiInX := math.Floor(time / (2 * math.Pi))
	x := (time-twoPiInX*2*math.Pi)/(1*math.Pi) - 1

	return []floatPoint{
		{x, cos},
		{x, -cos},
		{-x, cos},
		{-x, -cos},
		{cos, x},
		{-cos, x},
		{cos, -x},
		{-cos, -x},
	}
}

type figureState struct {
	nowFunc   figure
	changeMut sync.Mutex
	allFuncs  []figure
}

func newFigure(nowIdx int) *figureState {
	allFuncs := []figure{
		circle,
		itsAllCos,
		sinWithX,
	}
	return &figureState{
		allFuncs: allFuncs,
		nowFunc:  allFuncs[nowIdx],
	}
}

func (st *figureState) change(newIdx int) {
	st.changeMut.Lock()
	st.nowFunc = st.allFuncs[newIdx]
	st.changeMut.Unlock()
}

func (st *figureState) getCoords(time float64) []floatPoint {
	st.changeMut.Lock()
	points := st.nowFunc(time)
	st.changeMut.Unlock()
	return points
}
