package main

import (
	"fmt"
	"time"
)

func display(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}

func displaySequentialFirst(msg string, chhello chan<- int, chworld chan int) {
	fmt.Println(msg)
	chhello <- 1

	for {
		select {
		case <-chworld:
			fmt.Println(msg)
			chhello <- 1
		}
	}
}

func displaySequentialSecond(msg string, chhello chan int, chworld chan<- int) {
	for {
		select {
		case <-chhello:
			fmt.Println(msg)
			chworld <- 1
		}
	}
}

func main() {
	// go display("world")
	// display("hello")

	chhello, chworld := make(chan int), make(chan int)

	go displaySequentialFirst("hello", chhello, chworld)
	displaySequentialSecond("world", chhello, chworld)
}
