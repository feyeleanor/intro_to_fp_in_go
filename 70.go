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

func PrintErrorMessage() {
	if x := recover(); x != nil {
		Printf("no defined value for %v\n", x)
	}
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

func Each(s []string, f func(string)) {
	if len(s) > 0 {
		f(s[0])
		Each(s[1:], f)
	}
}

func OnPanic(e func()) func(func()) {
	return func(f func()) {
		defer e()
		f()
	}
}

var cache map[int] int = make(map[int] int)

func Factorial(n int) (r int) {
	if r = cache[n]; r == 0 {
		switch {
		case n < 0:
			panic(n)
		case n == 0:
			r = 1
		default:
			r = n * Factorial(n - 1)
		}
		cache[n] = r
	}
	return
}