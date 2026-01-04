package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func i2c8(x int) complex128 {
	floatx := float64(x)
	return complex(floatx, 0)
}

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

func testFftnaive() {
	N := 8
	list := make([]complex128, N)
	for i := range N {
		list[i] = complex(float64(i+1), 0) // Numbers from 1 to N
	}
	fftlist := fftnaive(list, N)
	fmt.Println(fftlist)
}

func testApply() {
	N := 8
	list := make([]float64, N)
	for i := range N {
		list[i] = float64(i + 1)
	}
	applysin := Apply(math.Sin, list)
	fmt.Println(applysin)
}
