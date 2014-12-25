package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

	z := 1.0
	tmp := x

	for {
		z = (z + x/z) / 2
		if z == tmp {
			break
		} else {
			tmp = z
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
