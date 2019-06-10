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

func main() {
	var wg sync.WaitGroup
	fmt.Println("Time now: ", time.Now())
	wg.Add(1)
	Announce("Done", time.Duration(2)*time.Second, &wg)
	fmt.Println("Time now: ", time.Now())
	wg.Wait()
}
