package main
import "os"
import "strconv"

func main() {
	a := MakeAccumulator()
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x)
	}
	os.Exit(a(0))
}

type Accumulator func(int) int

func MakeAccumulator() Accumulator {
	var y int
	return func(x int) int {
		y += x
		return y
	}
}