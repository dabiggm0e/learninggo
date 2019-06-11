package main

import (
	"fmt"
	"sync"
)

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func handle(queue chan *Request, results chan *Request, wg *sync.WaitGroup) {
	fmt.Println("[handle] Inside handle")

	for r := range queue {
		fmt.Println("[handle] Reading queue")
		r.resultChan <- r.f(r.args)
		close(r.resultChan)
		fmt.Println("[handle] Closing queue")
		results <- r
		wg.Done()
		fmt.Println("[handle] Done goroutine")
	}

	close(results)
	return
}

func main() {
	var wg sync.WaitGroup
	request := &Request{[]int{1, 2, 3, 4}, sum, make(chan int, 2)}
	queue := make(chan *Request, 100)
	results := make(chan *Request, 100)

	fmt.Printf("[main] Adding %v to the queue\n", *request)
	wg.Add(1)
	queue <- request
	fmt.Println("[main] closing the queue")
	close(queue)
	fmt.Println("[main] calling handle")
	go handle(queue, results, &wg)

	wg.Wait()
	fmt.Printf("len = %v, size = %v\n", len(results), cap(results))

	for r := range results {
		fmt.Printf("The sum of %v is %v\n", r.args, <-r.resultChan)
	}
}
