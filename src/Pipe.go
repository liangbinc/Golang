package main

import (
	"fmt"
	_ "io"
)

func init_pipe() chan int {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	fmt.Println(len(pipe))
	return pipe
}

func getPipe(pipe chan int) {

	for i := len(pipe); i > 0; i-- {
		fmt.Println(<-pipe)
	}
}

func main() {
	pipe := init_pipe()
	getPipe(pipe)
}
