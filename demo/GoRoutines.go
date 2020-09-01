package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wait = sync.WaitGroup{}

func concurrencyTestOne() {
	fmt.Printf("I am a Go Routine")
	wait.Done()
}

func routine1(p int) {
	fmt.Printf("I am function 1: %v\n", p)
	wait.Done()
}

func routine2(p int) {
	fmt.Printf("I am function 2: %v\n", p)
	wait.Done()
}

func goRoutineTest() {
	for i := 0; i < 10; i++ {
		wait.Add(2)

		go routine2(i)
		go routine1(i)

		wait.Wait()
		fmt.Printf("GAP\n")
	}
}

var mutex = sync.RWMutex{}
var counter = 0

func printHello() {
	fmt.Printf("I am %v\n", counter)
	mutex.RUnlock()
	wait.Done()
}

func writeCounter() {
	counter++
	mutex.Unlock()
	wait.Done()
}

func looping() {
	for p := 0; p < 10; p++ {
		wait.Add(2)
		mutex.RLock()
		go printHello()
		mutex.Lock()
		go writeCounter()
	}
}

func printAvailableThreads() {
	fmt.Printf("Available threads are %v", runtime.GOMAXPROCS(-1))
}
