package main
import "os"
import "strconv"

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		accumulate(x)
	}
	os.Exit(y)
}

var y int

func accumulate(x int) {
	y += x
}