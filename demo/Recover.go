package main

import (
	"fmt"
	"log"
)

//A panic function first executes defer and then panic message
//Recover anonymous function will let the parent method function after the anonymous function's execution

func showRecover() {
	fmt.Printf("START\n")
	panicker()
	fmt.Printf("STOP\n")
}

func panicker() {

	fmt.Printf("panicker() is about to panic\n")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error is ", err)
		}
	}()
	panic("I AM A PANIC MESSAGE")
	fmt.Printf("End of this function")
}
