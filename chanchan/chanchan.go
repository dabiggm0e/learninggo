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
	for n := range nums {
		sum += n
	}
	return sum
}

func handle(queue chan *Request, wg *sync.WaitGroup) {
	fmt.Println("Inside handle")

	for r := range queue {
		r.resultChan <- r.f(r.args)
		close(r.resultChan)
		fmt.Println("Done goroutine")
		wg.Done()
	}

	return
}

func main() {
	var wg sync.WaitGroup
	request := &Request{[]int{1, 2, 3, 4}, sum, make(chan int)}
	queue := make(chan *Request, 100)
	wg.Add(1)
	queue <- request

	close(queue)
	handle(queue, &wg)

	//	wg.Wait()
	fmt.Printf("len = %v, size = %v", len(queue), cap(queue))

	//fmt.Printf("The sum of %v is %v", r.args, <-r.resultChan)
	/*for r := range queue {
		fmt.Printf("The sum of %v is %v", r.args, <-r.resultChan)
	}*/
}
