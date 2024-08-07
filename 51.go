package main

import "os"

func main() {
	os.Exit(add(3, 4))
}

type Scalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func add[T Scalar](x, y T) T {
	return x + y
}
