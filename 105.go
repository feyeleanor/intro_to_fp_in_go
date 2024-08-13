package main

import "os"

type Function[T any] func(T) T
type Transformer[T any] func(Function[T]) Function[T]

type Recursive[T any] func(Recursive[T]) Function[T]

func (r Recursive[T]) Apply(f Transformer[T]) Function[T] {
	return f(r(r))
}

func Y[T any](f Transformer[T]) Function[T] {
	g := func(r Recursive[T]) Function[T] {
		return func(x T) T {
			return r.Apply(f)(x)
		}
	}
	return g(g)
}

func main() {
	factorial := Y(func(g Function[int]) Function[int] {
		return func(n int) int {
			switch {
			case n < 0:
				panic(n)
			case n < 2:
				return 1
			}
			return n * g(n-1)
		}
	})
	os.Exit(factorial(5))
}
