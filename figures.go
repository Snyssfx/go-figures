package main

import (
	"math"
)

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

func hypocycloid(time float64) []floatPoint {
	k := 5.5
	r := 1.0 / k
	x := r*(k - 1) * math.Cos(time) + r * math.Cos((k - 1) * time)
	y := r*(k - 1) * math.Sin(time) + r * math.Sin((k - 1) * time)

	return []floatPoint {
		{x, y},
	}
}