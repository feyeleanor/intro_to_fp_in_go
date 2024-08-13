package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := MakeAccumulator[int]()
	for i, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a.Add(x)
		a.Add(MakeAccumulator(i))
	}
	os.Exit(a.Int())
}

type Scalar interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Accumulator[T Scalar] func(T) T

func MakeAccumulator[T Scalar](s ...T) (a Accumulator[T]) {
	var y T
	a = func(x T) T {
		y += x
		return y
	}
	for _, v := range s {
		a.Add(v)
	}
	return
}

func (a Accumulator[T]) Int() int {
	return int(a(0))
}

func (a Accumulator[T]) Add(x any) {
	switch x := x.(type) {
	case T:
		a(x)
		fmt.Printf("1>a + x == a + %v == %v\n", x, a.Int())
	case Accumulator[T]:
		a(x(0))
		fmt.Printf("2>a + x == a + %v == %v\n", x(0), a.Int())
	}
}
