package main

import "fmt"

//TODO Learn tags from the video
type car struct {
	make              string
	model             string
	manufacturingYear uint
	contextData       map[string]string
}

func initializeCar(make string, model string, year uint) (car, int, string) {
	myCar := car{
		make: make, model: model, manufacturingYear: year,
		contextData: map[string]string{
			"Owner Name": "Mainaak",
			"Owner Type": "Business Owned",
		},
	}
	myCar.contextData["Owner Type"] = "New"
	fmt.Printf("The car is\n%v\n", myCar)
	fmt.Printf("The car's owner is %v", myCar.contextData["Owner Name"])
	return myCar, 10, "Argument"
}
