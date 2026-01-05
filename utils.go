package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

func Apply[T any](f func(T) T, list []T) []T {
	retval := make([]T, len(list))
	for i, v := range list {
		retval[i] = f(v)
	}
	return retval
}

func CompareEqualArrays(A, B []complex128, eps ...float64) bool {
	if len(A) != len(B) {
		return false
	}
	epsval := 1e-5
	if len(eps) > 0 {
		epsval = eps[0]
	}
	for i := range A {
		if cmplx.Abs(A[i]-B[i]) > epsval {
			return false
		}
	}
	return true
}

func Chop(x float64, thresh ...float64) float64 {
	threshval := 1e-15
	if len(thresh) > 0 {
		threshval = thresh[0]
	}

	if math.Abs(x-0.0) < threshval {
		x = 0.0
	}
	return x
}

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
