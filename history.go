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

	if len(h.points) == h.maxLen {
		h.points = h.points[1:]
	}

	h.points = append(h.points, points)
}
