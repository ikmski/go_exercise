package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var z, tmp complex128 = x, x
	for {
		z = z - (z*z*z-x)/(3.0*z*z)
		if cmplx.Abs(z-tmp) < 1e-10 {
			break
		}
		tmp = z
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
