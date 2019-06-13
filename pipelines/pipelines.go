package main

import (
	"fmt"
	"os"
	"time"
)

var done chan struct{}
var squarer chan int

func generateNaturals(naturals chan int) {
	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("Exiting.... Closing channel")
			close(naturals)

			return
		default:
			time.Sleep(time.Second * 1)
			naturals <- i
		}

	}
}

func generateSquarer(naturals chan int, squarer chan int) {
	for i := range naturals {

		squarer <- i * i
	}
	close(squarer)
}

func printSquarer(squarer chan int) {
	for i := range squarer {
		fmt.Println(i)
	}
}

func main() {

	var b = make([]byte, 1)
	var naturals = make(chan int)
	squarer = make(chan int)
	done = make(chan struct{})

	go func() {
		os.Stdin.Read(b)
		done <- struct{}{}
	}()

	go generateNaturals(naturals)
	go generateSquarer(naturals, squarer)
	printSquarer(squarer)

}
