package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		z = z - (z*z-x)/2*z
		if math.Abs(z-tmp) < 0.0001 {
			return z
		}
		tmp = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
