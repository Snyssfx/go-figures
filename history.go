package main

import (
	"reflect"
)

type history struct {
	points [][]intPoint
	maxLen int
}

func (h *history) add(points []intPoint) {
	if len(h.points) > 0 && reflect.DeepEqual(h.points[len(h.points)-1], points) {
		return
	}

	if len(h.points) >= h.maxLen {
		diff := len(h.points) - h.maxLen + 1
		h.points = h.points[diff:]
	}

	h.points = append(h.points, points)
}

func (h *history) incMaxLen() {
	h.maxLen++
}

func (h *history) decMaxLen() {
	h.maxLen--
	if h.maxLen <= 0 {
		h.maxLen = 1
	}
}
