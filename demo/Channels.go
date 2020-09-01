package main

import "fmt"

func channelFunction(){
	//This makes a channel
	ch := make(chan int, 50)
	wait.Add(2)
	//Telling that this extracts out of channel of int
	go func(ch <-chan int) {
		for i := range ch{
			fmt.Printf("%v\n", i)
		}
		wait.Done()
	}(ch)

	//This function puts into channel int
	go func(ch chan<- int ){
		ch <- 44
		ch <- 5
		ch <- 16
		close(ch)
		wait.Done()
	}(ch)
	wait.Wait()
}
