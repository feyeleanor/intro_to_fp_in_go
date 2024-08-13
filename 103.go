package main

import "os"

type Function func(any) any
type Transformer func(Function) Function

func Recurse(f Function) Function {
	return f(f).(Function)
}

func Y(g Transformer) Function {
	return Recurse(
		func(f any) any {
			return g(func(x any) any {
				return Recurse(f.(Function))(x)
			})
		})
}

func main() {
	factorial := Y(func(g Function) Function {
		return func(n any) any {
			if n, ok := n.(int); ok {
				switch {
				case n < 0:
					panic(n)
				case n < 2:
					return 1
				}
				return n * g(n-1).(int)
			}
			panic(n)
		}
	})
	os.Exit(factorial(5).(int))
}
