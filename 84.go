package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("out of stack space")
			fmt.Println(x)
			os.Exit(1)
		}
	}()
	main()
}
