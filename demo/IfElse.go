package main

import "fmt"

var a, b, c = 5, 10, 100

func ifConditions()  {
	if a > b || returnsTrue(){
		fmt.Printf("PASSED")
	} else {
		fmt.Printf("FAILED")
	}
}

func returnsTrue() bool {
	fmt.Printf("In the returnsTrue()")
	return true
}
