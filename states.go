package main

import (
	"sync"
)

type figureState struct {
	nowFunc   figure
	changeMut sync.Mutex
	allFuncs  []figure
	coef      float64
}

func newFigure(nowIdx int, startCoef float64) *figureState {
	allFuncs := []figure{
		circle,
		itsAllCos,
		sinWithX,
		epicycloid,
		hypocycloid,
	}
	return &figureState{
		allFuncs: allFuncs,
		nowFunc:  allFuncs[nowIdx],
		coef:     startCoef,
	}
}

func (st *figureState) change(newIdx int) {
	if 0 <= newIdx && newIdx < len(st.allFuncs) {
		st.changeMut.Lock()
		st.nowFunc = st.allFuncs[newIdx]
		st.changeMut.Unlock()
	}
}

func (st *figureState) getCoords(time float64) []floatPoint {
	st.changeMut.Lock()
	points := st.nowFunc(time, st.coef)
	st.changeMut.Unlock()
	return points
}

func (st *figureState) incCoef() {
	st.coef += 0.1
}

func (st *figureState) decCoef() {
	st.coef -= 0.1
	if st.coef <= 0.05 {
		st.coef = 0.1
	}
}
