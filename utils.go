package main

func Apply[T any](f func(T) T, list []T) []T {
	retval := make([]T, len(list))
	for i, v := range list {
		retval[i] = f(v)
	}
	return retval
}
