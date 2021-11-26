package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	workerCount := 1
	wg := &sync.WaitGroup{}
	wg.Add(workerCount)
	workerQueue := make(chan map[int]int, workerCount)

	for i := 0; i < workerCount; i++ {
		go testFunc(wg, workerQueue)
	}

	elapsed := time.Since(start)
	log.Printf("Test took %s", elapsed)
	wg.Wait()
}

func testFunc(wg *sync.WaitGroup, c <-chan map[int]int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		testRef(wg, c)
	}
}

func testRef(wg *sync.WaitGroup, c <-chan map[int]int) {
	i := 10
	i *= i * i * i * i * i * i
}
