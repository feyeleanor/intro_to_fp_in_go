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
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}