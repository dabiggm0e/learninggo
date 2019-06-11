package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Hello from service1"
}

func service2(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Hello from service2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	c1 := make(chan string)
	c2 := make(chan string)

	go service1(c1)
	go service2(c2)

	select {
	case res := <-c1:
		fmt.Println(res, time.Since(start))

	case res := <-c2:
		fmt.Println(res, time.Since(start))

	case <-time.After(4 * time.Second):
		fmt.Println("No response received", time.Since(start))

		/*default:
		fmt.Println("No response received", time.Since(start))*/
	}

	fmt.Println("main() ended", time.Since(start))
}
