package main

import (
	"fmt"
	"runtime"
)

func main() {
	numberOfRequests := 100
	maxWorkerNumber := 5

	queueChan := make(chan int, numberOfRequests)
	doneChan := make(chan int)

	for i := 0; i < maxWorkerNumber; i++ {
		go func(name string) {
			for v := range queueChan {
				crawl(name, v)
			}
			fmt.Printf("%s is done\n", name)
			doneChan <- 1
		}(fmt.Sprintf("%d", i))
	}

	for i := 0; i < numberOfRequests; i++ {
		queueChan <- i
	}
	close(queueChan)

	for i := 0; i < maxWorkerNumber; i++ {
		<-doneChan
	}
}
func crawl(name string, v int) {
	runtime.Gosched() // switch other process
	fmt.Printf("Worker %s is crawling: %d\n", name, v)
}
