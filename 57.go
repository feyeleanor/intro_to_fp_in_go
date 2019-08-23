package main
import . "fmt"
import "os"
import "strconv"

func main() {
	for _, v := range os.Args[1:] {
		OnPanicFor(UseNumericParam(v, func(i int) {
			Printf("%v!: %v\n", i, Factorial(i))
		}), PrintErrorMessage)
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

func OnPanicFor(f, e func()) {
	defer e()
	f()
}

func PrintErrorMessage() {
	if x := recover(); x != nil {
		Printf("no defined value for %v\n", x)
	}
}

func Factorial(n int) (r int) {
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}