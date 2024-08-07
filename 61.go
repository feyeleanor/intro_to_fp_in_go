package main

import (
	"os"
	"strconv"
)

func main() {
	a := MakeAccumulator[int]()
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x)
	}
	os.Exit(a.Int())
}

type Scalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Accumulator[T Scalar] func(T) T

func MakeAccumulator[T Scalar]() Accumulator[T] {
	var y T
	return func(x T) T {
		y += x
		return y
	}
}

func (a Accumulator[T]) Int() int {
	return int(a(0))
}
