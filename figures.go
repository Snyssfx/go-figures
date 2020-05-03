package main

import (
	"math"
)

// return x, y from -1 to 1
type figure func(time, coef float64) []floatPoint

func circle(time, coef float64) []floatPoint {
	x := math.Cos(time) * coef
	y := math.Sin(time) * coef
	return []floatPoint{{x, y}}
}

func sinWithX(time, coef float64) []floatPoint {
	y := math.Sin(coef * time)

	twoPiInX := math.Floor(time / (2 * math.Pi))
	x := (time-twoPiInX*2*math.Pi)/(1*math.Pi) - 1

	return []floatPoint{
		{x, y},
		{x, 0},
	}
}

func itsAllCos(time, coef float64) []floatPoint {
	cos := math.Cos(coef * time)

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

func epicycloid(time, coef float64) []floatPoint {
	k := coef
	if k < 1.1 {
		k = 1.1
	}
	r := 1.0 / k
	x := r*(k-1)*math.Cos(time) + r*math.Cos((k-1)*time)
	y := r*(k-1)*math.Sin(time) + r*math.Sin((k-1)*time)

	return []floatPoint{
		{x, y},
	}
}

func hypocycloid(time, coef float64) []floatPoint {
	k := coef
	if k < 1.1 {
		k = 1.1
	}
	r := 1.0 / k
	x := r*(k-1)*math.Cos(time) + r*math.Cos((k-1)*time)
	y := r*(k-1)*math.Sin(time) - r*math.Sin((k-1)*time)

	return []floatPoint{
		{x, y},
	}
}
