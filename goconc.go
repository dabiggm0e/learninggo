package main

import (
	"fmt"
	"sync"
	"time"
)

//Announce announces a message to standard output after a certain amount of time
func Announce(message string, delay time.Duration, wg *sync.WaitGroup) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message, time.Now())
		wg.Done()
	}()
}

func greet(roc <-chan string) {
	fmt.Println(<-roc)
}

func main() {
	fmt.Println("Greeting")
	ch := make(chan string)
	go greet(ch)
	ch <- "Moe"

	var wg sync.WaitGroup
	fmt.Println("Testing sync.WaitGroup")
	fmt.Println("Testing ...")
	fmt.Println("Time now: ", time.Now())
	wg.Add(1)
	Announce("Done", time.Duration(2)*time.Second, &wg) // COMBAK:
	fmt.Println("Time now: ")
	wg.Wait()

}
