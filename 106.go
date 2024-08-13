package main

import "fmt"

type Function[T any] func(T) T
type Transformer[T any] func(Function[T]) Function[T]

type Recursive[T any] func(Recursive[T]) Function[T]

func (r Recursive[T]) Apply(f Transformer[T]) Function[T] {
	return f(r(r))
}

func Y[T comparable](f Transformer[T]) Function[T] {
	memo := make(map[T]T)
	g := func(r Recursive[T]) Function[T] {
		return func(x T) T {
			if _, ok := memo[x]; !ok {
				memo[x] = r.Apply(f)(x)
				fmt.Printf("Y: setting memo[%v] = %v\n", x, memo[x])
			}
			return memo[x]
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

	fmt.Printf("%v! = %v\n", 7, factorial(7))
	fmt.Printf("%v! = %v\n", 8, factorial(8))
	fmt.Printf("%v! = %v\n", 5, factorial(5))
}
