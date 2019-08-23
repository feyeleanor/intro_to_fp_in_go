package main
import . "fmt"
import "os"
import "strconv"

func main() {
	var errors int
	for _, v := range os.Args[1:] {
		SafeExecute(
			func(i int) {
				Printf("%v!: %v\n", i, Factorial(i))
			},
			func() {
				if x := recover(); x != nil {
					Printf("no defined value for %v\n", x)
					errors++
				}
			},
		)(v)
	}
	os.Exit(errors)
}

func SafeExecute(f func(int), e func()) func(string) {
	return func(v string) {
		defer e()

		if x, e := strconv.Atoi(v); e == nil {
			f(x)
		} else {
			panic(v)
		}
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