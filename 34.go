package main
import "os"
import "strconv"

func main() {
	os.Exit(add(arg(0), arg(1)))
}

func arg(n int) (r int) {
	r, _ = strconv.Atoi(os.Args[n + 1])
	return
}

func add(x, y int) int {
	return x + y
}