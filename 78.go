package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var errors int
	f := func(i int) {
		Printf("%v!: %v\n", i, Factorial(i))
	}
	for _, v := range os.Args[1:] {
		if !ForValidValues(f)(v) {
			errors++
		}
	}
	os.Exit(errors)
}

func ForValidValues[T Integer](f func(T)) func(string) bool {
	return func(v string) (r bool) {
		defer func() {
			if x := recover(); x != nil {
				Printf("no defined value for %v\n", x)
			}
		}()

		if x, e := strconv.Atoi(v); e == nil {
			f(T(x))
			r = true
		} else {
			panic(v)
		}
		return
	}
}

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
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
