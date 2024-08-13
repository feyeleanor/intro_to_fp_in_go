package main

import "os"

type Function func(any) any

func Recurse(f any) Function {
	return f.(func(any) any)(f).(func(any) any)
}

func Y(g Function) Function {
	return Recurse(
		func(f any) any {
			return g(func(x any) any {
				return Recurse(f)(x)
			})
		})
}

func main() {
	factorial := Y(func(g any) any {
		return func(n any) any {
			if n, ok := n.(int); ok {
				switch {
				case n < 0:
					panic(n)
				case n < 2:
					return 1
				}
				return n * g.(func(any) any)(n-1).(int)
			}
			panic(n)
		}
	})
	os.Exit(factorial(5).(int))
}
