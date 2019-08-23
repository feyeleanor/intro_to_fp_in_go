package main
import . "fmt"
import "os"
import "strconv"

func main() {
	f := MakeFactorial()
	Each(os.Args[1:], func(v string) {
		OnPanic(PrintErrorMessage)(
			UseNumericParam(v, func(i int) {
				Printf("%v!: %v\n", i, f(i))
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

func MakeFactorial() (f func(int) int) {
	c := make(map[int] int)
	return func(n int) (r int) {
		if r = c[n]; r == 0 {
			switch {
			case n < 0:
				panic(n)
			case n == 0:
				r = 1
			default:
				r = n * f(n - 1)
			}
			c[n] = r
		}
		return
	}
}