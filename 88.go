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
	for _, v := range os.Args[1:] {
		if x, e := strconv.Atoi(v); e == nil {
			Printf("%v!: %v\n", x, Factorial(x))
		} else {
			panic(v)
		}
	}
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
