package main
import . "fmt"
import "os"
import "strconv"

func main() {
	Each(os.Args[1:], func(v string) {
		OnPanic(PrintErrorMessage)(
			UseNumericParam(v, func(i int) {
				Printf("%v!: %v\n", i, Factorial(i))
		}))
	})
}

func UseNumericParam(v string, f func(i int)) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			f(x)
		} else {
			panic(v)
		}
	}
}

func OnPanic(e func()) func(func()) {
	return func(f func()) {
		defer e()
		f()
	}
}

func Each(s []string, f func(string)) {
	for _, v := range s {
		f(v)
	}
}

func PrintErrorMessage() {
	if x := recover(); x != nil {
		Printf("no defined value for %v\n", x)
	}
}

func Factorial(n int) (r int) {
	switch {
	case n < 0:
		panic(n)
	case n == 0:
		r = 1
	default:
		r = n * Factorial(n - 1)
	}
	return
}