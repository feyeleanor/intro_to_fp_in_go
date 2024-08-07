package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	printErrors := IfPanics(PrintErrorMessage)
	for _, v := range os.Args[1:] {
		printErrors(
			ValidInteger(v, func(i int) {
				Printf("%v!: %v\n", i, Factorial(i))
			}))
	}
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

func PrintErrorMessage() {
	if x := recover(); x != nil {
		Printf("no defined value for %v\n", x)
	}
}

func Factorial[T Integer](n T) (r T) {
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}
