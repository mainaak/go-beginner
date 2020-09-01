package main

import "fmt"

func execFunction() {

	name := "Amit"
	age, year := 27, 1992
	pointerFunction(&name, &age, &year)
	fmt.Printf("Updated values are %v, %v and %v\n", name, age, year)
}

func pointerFunction(name *string, age, year *int) {

	fmt.Printf("Original values are %v, %v, %v\n", *name, *age, *year)
	*name = "Mainaak"
	*age, *year = 25, 1994
	fmt.Printf("%v was born on %v and is %v years old now\n", *name, *age, *year)
}
