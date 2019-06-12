// source: https://golang.org/doc/codewalk/sharemem/
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2
	pollInterval   = 60 * time.Second
	statusInterval = 10 * time.Second
	errTimeout     = 10 * time.Second
)

var urls = []string{
	"http://www.google.com",
	"http://golang.org/",
	"http://blog.golang.org/",
}

//State represents the last known status of a URL
type State struct {
	url    string
	status string
}

func StateMonitor(statusInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(statusInterval)

	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()

	return updates
}

func logState(s map[string]string) {
	fmt.Println("Current state:")
	for k, v := range s {
		log.Printf("%s: %s\n", k, v)
	}
}

type Resource struct {
	url      string
	errCount int
}

func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		r.errCount++
		log.Println("Error: ", r.url, err)
		return err.Error()
	}

	r.errCount = 0
	return resp.Status
}

func (r *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func main() {
	pending, complete := make(chan *Resource), make(chan *Resource)

	status := StateMonitor(statusInterval)

	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}
}
