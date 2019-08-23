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
	x, _ := strconv.Atoi(os.Args[1])
	Printf("%v!: %v\n", x, Factorial(x))
}

func Factorial(n int) (r int) {
	switch {
	case n < 0:
		panic(n)
	case n == 0:
		r = 1
	default:
		r = 1
		for ; n > 0; n-- {
			r *= n
		}
	}
	return
}