package main

type history struct {
	points [][]point
	length int
}

func (h *history) add(points []point) {
	if len(h.points) == h.length {
		h.points = h.points[1:]
	}
	h.points = append(h.points, points)
}
