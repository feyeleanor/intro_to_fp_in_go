package main

import (
	"fmt"
	"math"
)

func main() {
	defer func() {
		recover()
	}()
	for i := range math.MaxInt {
		fmt.Printf("%v: %v\n", i, []int{0, 2, 4, 6, 8}[i])
	}
}
