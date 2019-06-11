package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
