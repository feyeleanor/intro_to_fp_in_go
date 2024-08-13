package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f := MakeFactorial[int]()
	printErrors := IfPanics(PrintErrorMessage)
	Each(os.Args[1:], func(v string) {
		printErrors(
			ValidInteger(v, func(i int) {
				fmt.Printf("%v!: %v\n", i, f(i))
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
	if len(s) > 0 {
		f(s[0])
		Each(s[1:], f)
	}
}

func PrintErrorMessage() {
	if x := recover(); x != nil {
		fmt.Printf("no defined value for %v\n", x)
	}
}

func MakeFactorial[T Integer]() (f func(T) T) {
	c := make(map[T]T)
	return func(n T) (r T) {
		if r = c[n]; r == 0 {
			switch {
			case n < 0:
				panic(n)
			case n == 0:
				r = 1
			default:
				r = n * f(n-1)
			}
			c[n] = r
		}
		return
	}
}
