package main
import . "fmt"

func main() {
	c := make(chan int)
	go sequence(c)
	print_channel(c)
}

func sequence(c chan int) {
	for i := 0; ; i++ {
		c <- i * 2
	}
}

func print_channel(c chan int) {
	for i := 0; i < 5; i++ {
		Printf("%v: %v\n", i, <- c)
	}
	return
}