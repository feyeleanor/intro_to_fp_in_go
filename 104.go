package main

import "os"

type Function func(any) any
type Transformer func(Function) Function

type Recursive func(Recursive) Function

func (r Recursive) Apply(f Transformer) Function {
	return f(r(r))
}

func Y(f Transformer) Function {
	g := func(r Recursive) Function {
		return func(x any) any {
			return r.Apply(f)(x)
		}
	}
	return g(g)
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
