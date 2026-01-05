package main

import (
	"math"
	"math/cmplx"
	"sync"
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
	// defer Timer("FFTnaive")() //dumb, because this function is recursive
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

func fftgov1(a []complex128, ch chan []complex128) { // will lead to race conditions, but we ignore for now
	n := len(a)
	if n == 2 {
		ch <- []complex128{a[0] + a[1], a[0] - a[1]}
		return
	}

	halfsize := n / 2
	u := make([]complex128, halfsize)
	v := make([]complex128, halfsize)
	for j := range halfsize {
		u[j] = a[2*j]
		v[j] = a[2*j+1]
	}
	chi := make(chan []complex128)
	go fftgov1(u, chi)
	U := <-chi
	go fftgov1(v, chi)
	V := <-chi
	argn := 2.0 * math.Pi / float64(n)
	omegan := complex(math.Cos(argn), math.Sin(argn))
	omega := 1.0 + 0i
	A := make([]complex128, n)

	for j := range halfsize {
		A[j] = U[j] + omega*V[j]
		A[j+halfsize] = U[j] - omega*V[j]
		omega = omega * omegan
	}
	ch <- A
}

func FFTgov1(a []complex128) []complex128 {
	// defer Timer("FFTgo")()
	ch := make(chan []complex128)
	go fftgov1(a, ch)
	A := <-ch
	return A
}

func fftgo(a []complex128, A *[]complex128) { // A is the fourier transform; still need to use a mutex I suppose
	n := len(a)
	if n == 2 {
		(*A)[0] = a[0] + a[1]
		(*A)[1] = a[0] - a[1]
		return
	}

	halfsize := n / 2
	u := make([]complex128, halfsize)
	v := make([]complex128, halfsize)
	for j := range halfsize {
		u[j] = a[2*j]
		v[j] = a[2*j+1]
	}
	U := make([]complex128, halfsize)
	V := make([]complex128, halfsize)
	var wc sync.WaitGroup
	wc.Go(func() { fftgo(u, &U) })
	wc.Go(func() { fftgo(v, &V) })
	// go fftgo(u, &U)
	// go fftgo(v, &V)
	argn := 2.0 * math.Pi / float64(n)
	omegan := complex(math.Cos(argn), math.Sin(argn))
	omega := 1.0 + 0i
	wc.Wait()

	for j := range halfsize {
		(*A)[j] = U[j] + omega*V[j]
		(*A)[j+halfsize] = U[j] - omega*V[j]
		omega = omega * omegan
	}
}

func FFTgo(a []complex128) []complex128 {
	A := make([]complex128, len(a))
	fftgo(a, &A)
	return A
}
