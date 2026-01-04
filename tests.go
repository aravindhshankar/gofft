package main

import (
	"fmt"
	"math"
)

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
