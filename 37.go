package main
import "os"
import "strconv"

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a.Add(x)
	}
	os.Exit(int(a))
}

var a Accumulator

type Accumulator int

func (a *Accumulator) Add(y int) {
	*a += Accumulator(y)
}