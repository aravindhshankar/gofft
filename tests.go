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
