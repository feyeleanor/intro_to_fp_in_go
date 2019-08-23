package main
import "os"
import "strconv"

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x)
	}
	os.Exit(a(0))
}

var a = MakeAccumulator()

func MakeAccumulator() func(int) int {
	var y int
	return func(x int) int {
		y += x
		return y
	}
}