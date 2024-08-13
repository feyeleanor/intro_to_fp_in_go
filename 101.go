package main

import "os"

// Y = λf.(λx.f(x x))(λx.f(x x))

func Y(g func(any) any) func(any) any {
	return func(f any) func(any) any {
		return f.(func(any) any)(f).(func(any) any)
	}(func(f any) any {
		return g(func(x any) any {
			return f.(func(any) any)(f).(func(any) any)(x)
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
