package main
import . "fmt"
import "os"
import "strconv"

func main() {
	defer func() {
		if x := recover(); x != nil {
			Println("no factorial")
		}
	}()
	for _, v := range os.Args[1:] {
		if x, e := strconv.Atoi(v); e != nil {
			Printf("%v!: %v\n", x, Factorial(x))
		} else {
			panic(v)
		}
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