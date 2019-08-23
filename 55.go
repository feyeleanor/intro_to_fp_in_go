package main
import . "fmt"
import "os"
import "strconv"

func main() {
	var errors int
	f := func(i int) {
		Printf("%v!: %v\n", i, Factorial(i))
	}
	for _, v := range os.Args[1:] {
		if !SafeExecute(f)(v) {
			errors++
		}
	}
	os.Exit(errors)
}

func SafeExecute(f func(int)) func(string) bool {
	return func(v string) (r bool) {
		defer func() {
			if x := recover(); x != nil {
				Printf("no defined value for %v\n", x)
			}
		}()

		if x, e := strconv.Atoi(v); e == nil {
			f(x)
			r = true
		} else {
			panic(v)
		}
		return
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