package main

import (
	"math"
	"math/cmplx"
)

func DFTnaive(a []complex128, size int) []complex128 {
	A := make([]complex128, size)
	for i := range size {
		for j := range size {
			fftfactor := cmplx.Exp(2.0i * math.Pi * i2c8(i*j) / i2c8(size))
			A[i] += fftfactor * a[j]
		}
	}
	return A
}

func FFTnaive(a []complex128) []complex128 {
	n := len(a)
	if n == 1 {
		return a
	}
	if n == 2 {
		return []complex128{a[0] + a[1], a[0] - a[1]}
	}

	halfsize := n / 2
	u := make([]complex128, halfsize)
	v := make([]complex128, halfsize)
	for j := range halfsize {
		u[j] = a[2*j]
		v[j] = a[2*j+1]
	}
	U := FFTnaive(u)
	V := FFTnaive(v)
	argn := 2.0 * math.Pi / float64(n)
	omegan := complex(math.Cos(argn), math.Sin(argn))
	omega := 1.0 + 0i
	A := make([]complex128, n)

	for j := range halfsize {
		A[j] = U[j] + omega*V[j]
		A[j+halfsize] = U[j] - omega*V[j]
		omega = omega * omegan
	}
	return A
}

func main() {
	// NdftNfft()
	testChop()
}
