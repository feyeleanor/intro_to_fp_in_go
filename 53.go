package main

import (
	"os"
	"strconv"
)

func main() {
	os.Exit(add(arg(0), arg(1)))
}

func arg(n int) (r int) {
	r, _ = strconv.Atoi(os.Args[n+1])
	return
}

type Scalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func add[T Scalar](x, y T) T {
	return x + y
}
