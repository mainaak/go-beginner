package main

import "fmt"

func mainExec() {
	c := automobile{
		model: "530D",
		make:  "BMW",
		year:  2001,
	}
	c.executeMe()
	d := automobile{make: "Audi"}
	d.executeMe()
}

type automobile struct {
	model, make string
	year        int
}

func (c automobile) executeMe() {
	fmt.Printf("Car:\t Model: %v:\tMake: %v:\tYear: %v\n", c.model, c.make, c.year)
}
