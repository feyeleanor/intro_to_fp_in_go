package main
import . "fmt"
import "os"
import "strconv"

func main() {
	for _, v := range os.Args[1:] {
		SafeExecute(func(i int) {
			Printf("%v!: %v\n", i, Factorial(i))
		})(v)
	}
}

func SafeExecute(f func(int)) func(string) {
	return func(v string) {
		defer func() {
			if x := recover(); x != nil {
				Printf("no defined value for %v\n", x)
			}
		}()

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