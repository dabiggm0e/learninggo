package main

import (
	"fmt"
	"sync"
)

func getInputChan() <-chan int {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	inputChan := make(chan int, 100)

	go func() {
		for n := range numbers {
			inputChan <- n
		}

		close(inputChan)
	}()

	return inputChan
}

func getSquaredChan(inputChan <-chan int) <-chan int {

	outputChan := make(chan int, 100)
	go func() {
		for n := range inputChan {
			outputChan <- n * n
		}

		close(outputChan)
	}()

	return outputChan
}

func merge(outputChans ...<-chan int) <-chan int {

	merged := make(chan int, 100)
	var wg sync.WaitGroup

	wg.Add(len(outputChans))

	output := func(sc <-chan int) {
		for sqr := range sc {
			merged <- sqr
		}

		wg.Done()
	}

	for _, outputChan := range outputChans {
		go output(outputChan)
	}

	//go func() {
	wg.Wait()
	close(merged)
	//}()

	return merged
}

func main() {

	chanInputNums := getInputChan()

	chanOutputSqr1 := getSquaredChan(chanInputNums)
	chanOutputSqr2 := getSquaredChan(chanInputNums)

	merged := merge(chanOutputSqr1, chanOutputSqr2)

	var sqrSum int
	for sqr := range merged {
		sqrSum += sqr
	}

	fmt.Println("The sum of all square roots between 0-9 is", sqrSum)
}
