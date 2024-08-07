package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			Println("no factorial")
		}
	}()
	x, _ := strconv.Atoi(os.Args[1])
	Printf("%v!: %v\n", x, Factorial(x))
}

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Factorial[T Integer](n T) (r T) {
	switch {
	case n < 0:
		panic(n)
	case n == 0:
		r = 1
	default:
		r = n * Factorial(n-1)
	}
	return
}
