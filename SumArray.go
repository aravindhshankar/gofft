package main

import (
	"fmt"
)

func sum(ch chan int, arr []int) {
	total := 0
	for _, val := range arr {
		total = total + val
	}
	ch <- total
}

func main2() {
	var total int
	c := make(chan int, 2)
	N := 10
	list := make([]int, N)
	for i := range N {
		list[i] = i + 1 // Numbers from 1 to N
	}

	go sum(c, list[0:N/2])
	go sum(c, list[N/2:N])
	total = <-c
	total += <-c
	// can be refactored to push to another channel and sum that channel
	fmt.Println(total)
}
