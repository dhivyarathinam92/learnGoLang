package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
