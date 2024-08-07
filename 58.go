package main

import (
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x)
	}
	os.Exit(a(0))
}

var a = MakeAccumulator[int]()

type Scalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func MakeAccumulator[T Scalar]() func(T) T {
	var y T
	return func(x T) T {
		y += x
		return y
	}
}
