package main

import "fmt"

var hmm func(number int)

func anonymousFunc() {

	something := func() {
		fmt.Printf("I am something()\n")
	}

	func() {
		fmt.Printf("I get executed first\n")
	}()

	something()

	hmm = func(number int){
		fmt.Printf("Lewis Hamilton is %v\n", number)
	}

	hmm(44)
}
