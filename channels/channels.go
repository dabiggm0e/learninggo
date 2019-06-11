package main

import (
	"fmt"
	"sync"
	"time"
)

var index int

func sqrWorker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v]: Sending results by worker %v\n", instance, instance)
		results <- num * num
	}
	wg.Done()
}

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	index += 1
	m.Unlock()

	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	fmt.Println("Writing 5 tasks")
	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}
	close(tasks)
	fmt.Println("5 tasks written")

	fmt.Println("Starting workers")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sqrWorker(tasks, results, &wg, i)
	}
	fmt.Println("3 workers started")
	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Printf("[main] Result %v: %v\n", i, <-results)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Index value is", index)

}
