package main

import (
	"sync"
)

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
		hypocycloid,
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