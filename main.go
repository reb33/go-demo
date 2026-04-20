package main

import "math"

func main() {
	const IMT_POWER = 2
	var userHeight = 1.8
	var userWeight= 100

	var imt = float64(userWeight) / math.Pow(userHeight, IMT_POWER)
	println("IMT:", imt)
}
