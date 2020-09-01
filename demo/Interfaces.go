package main

import "fmt"

//This is like a class file
type iAmBlank struct {
}

//This is an interface with a method
type iAmGod interface {
	xxxx(number int) int
}

//This is implementation of the method by the struct which implemented it
func (a iAmBlank) xxxx(number int) int {
	fmt.Printf("I got number %v and I returned %v\n", number, number*2)
	return number * 2
}

func iAmExecutable(number int) {

	//This is where we say that iAmBlank implements iAmGod
	var variableOne iAmGod = iAmBlank{}
	iAmDoubleTheNumber := variableOne.xxxx(number)
	fmt.Printf("Double of %v is %v", number, iAmDoubleTheNumber)
}
