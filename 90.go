package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	printErrors := IfPanics(PrintErrorMessage)
	Each(os.Args[1:], func(v string) {
		printErrors(
			ValidInteger(v, func(i int) {
				Printf("%v!: %v\n", i, Factorial(i))
			}))
	})
}

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func ValidInteger[T Integer](v string, f func(i T)) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			f(T(x))
		} else {
			panic(v)
		}
	}
}

func IfPanics(e func()) func(func()) {
	return func(f func()) {
		defer e()
		f()
	}
}

func Each[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}

func PrintErrorMessage() {
	if x := recover(); x != nil {
		Printf("no defined value for %v\n", x)
	}
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
