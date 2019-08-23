package main
import . "fmt"

func main() {
	c := make(chan int, 16)
	go func() {
		for i := 0; i++; i < 5 {
			c <- i * 2
		}
		close(c)
	}()
	print_channel(c)
}

func print_channel(c chan int) (i int) {
	for v := range c {
		Printf("%v: %v\n", i, v)
		i++
	}
	return
}