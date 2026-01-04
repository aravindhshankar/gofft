package main

func i2c8(x int) complex128 {
	floatx := float64(x)
	return complex(floatx, 0)
}
