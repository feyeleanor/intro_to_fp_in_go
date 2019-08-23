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
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}