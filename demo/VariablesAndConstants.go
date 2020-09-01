package main

import (
	"fmt"
	"strconv"
)

func learningOne() {

	x, y, z := 10, "Mainaak", 32.00

	var p, o, i string
	p = "EmptyString"
	o, i = "o", "i\n"
	fmt.Printf("the variables are: %v %v %v %v %v %v", x, y, z, p, o, i)

	var inBlock int = 50
	fmt.Printf("Number %v is of type %T \n", outScope, outScope)
	fmt.Printf("Number %v is of type %T \n", inBlock, inBlock)
	//	Typecasting variables
	var createString = strconv.Itoa(inBlock)
	fmt.Println("String createString is", createString)
	//	Typecasting between float and int
	valueInteger := 50
	fmt.Printf("Value of valueInteger = %v and its type is %T", valueInteger, valueInteger)
	var valueFloat float32
	valueFloat = float32(valueInteger)
	fmt.Printf("\nValue of valueFloat = %v and its type is %T", valueFloat, valueFloat)
}

const (
	one = iota
	two
)

func checkingOutConstants() {
	const mainaakOne = 25
	const mainaakTwo = 52
	fmt.Printf("\nValue 1 is %v and Value 2 is %v", mainaakOne, mainaakTwo)
	fmt.Printf("\nValue one is %v and Value two is %v", one, two)
}
