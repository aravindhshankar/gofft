package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func testDFTnaive() {
	N := 8
	list := make([]complex128, N)
	for i := range N {
		list[i] = complex(float64(i+1), 0) // Numbers from 1 to N
	}
	fftlist := DFTnaive(list, N)
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

func chopset(a float64) float64 {
	return Chop(a)
}

func testChop() {
	x := 1.5
	y := 2.1e-16
	z := []float64{2.3, 6, 1.1e-8, 2.1e-16}
	x = Chop(x)
	y = Chop(y)
	z = Apply(chopset, z)
	fmt.Println("Chopped x = ", x)
	fmt.Println("Chopped y = ", y)
	fmt.Println("Chopped z = ", z)
}

func NdftNfft() {
	N := 8
	list := make([]complex128, N)
	for i := range N {
		list[i] = complex(math.Pi*float64(i)/float64(N), 0)
	}
	list = Apply(cmplx.Sin, list)
	dftlist := DFTnaive(list, N)
	fmt.Println("dftlist = ", dftlist)

	fftlist := FFTnaive(list)
	fmt.Println("fftlist = ", fftlist)

	ok := CompareEqualArrays(dftlist, fftlist)
	fmt.Println(ok)
}
