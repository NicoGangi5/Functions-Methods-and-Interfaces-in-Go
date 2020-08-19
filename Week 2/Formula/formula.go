package main

import (
	"fmt"
	"strconv"
)

var acceleration, iVelocity, iDisplacement, time float64

func main()  {
	acceleration, iVelocity, iDisplacement, time = inputs()
	fmt.Println()
	fmt.Println("Acceleration:", acceleration)
	fmt.Println("Initial velocity:", iVelocity)
	fmt.Println("Initial displacement:", iDisplacement)
	fmt.Println("Time:", time)
	fmt.Println()

	fn := GenDisplaceFn(acceleration, iVelocity, iDisplacement)
	fmt.Println(fn(time))
}

func inputs() (float64, float64, float64, float64) {
	var acceleration, iVelocity, iDisplacement, time string

	fmt.Print("Enter the acceleration: ")
	fmt.Scan(&acceleration)
	acc, _ := strconv.ParseFloat(acceleration, 64)

	fmt.Print("Enter the initial velocity: ")
	fmt.Scan(&iVelocity)
	iVel, _ := strconv.ParseFloat(iVelocity, 64)

	fmt.Print("Enter the initial displacement: ")
	fmt.Scan(&iDisplacement)
	iDis, _ := strconv.ParseFloat(iDisplacement, 64)

	fmt.Print("Enter the time: ")
	fmt.Scan(&time)
	t, _ := strconv.ParseFloat(time, 64)

	return acc, iVel, iDis, t
}

func GenDisplaceFn(acc, iVel, iDis float64) func(float64) float64 {
	fn := func(t float64) float64{
		return ((0.5 * acc * t * t) + (iVel * t) + iDis)
	}
	return fn
}