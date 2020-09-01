package main

import "fmt"

func checkingSwitchCases(argument string)string {

	switch argument {
	case "One","Three","Five":
		return "The number is 1"
	case "Two":
		return "The number is 2"
	default:
		return "Not 1 or 2"
	}

}

func taglessSwitch(){
	i := 1
	switch {
	case i <= 10:
		fmt.Printf("%v is lesser than or equal to 10", i)
		fallthrough //If we use fallthrough, it will not break
	case i <= 20:
		fmt.Printf("The number %v is greater than 10 but lesser than 20", i)
	}
}

func typeSwitch(jack interface{}){
	switch jack.(type) {
	case int:
		fmt.Printf("This is of type integer")
	case string:
		fmt.Printf("This is of type string")
	case bool:
		fmt.Printf("This is of type boolena")
	default:
		fmt.Printf("Type is %T", jack)
	}
}
