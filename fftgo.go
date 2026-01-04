package main

import (
	"math"
	"math/cmplx"
)

func fftnaive(a []complex128, size int) []complex128 {
	A := make([]complex128, size)
	for i := range size {
		for j := range size {
			fftfactor := cmplx.Exp(2.0i * math.Pi * i2c8(i*j) / i2c8(size))
			A[i] = fftfactor * a[j]
		}
	}
	return A
}

func main() {
	testApply()
}
