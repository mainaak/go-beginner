package main

import "fmt"

func slices() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	afterFive := slice[4:]
	ignoreLastFive := slice[:len(slice)-5]
	betweenThreeAndSix := slice[3:6]
	fmt.Printf("%v\n", afterFive)
	fmt.Printf("%v\n", ignoreLastFive)
	fmt.Printf("%v\n", betweenThreeAndSix)
}

func arrays() {
	//array := [5]int{2,3}
	array := [...]int{5, 77, 44, 33}

	arrayCopy := &array
	arrayCopy[0] = 19

	var jack [2]string
	jack[0] = "Mainaak"
	jack[1] = "Sharma"

	//array[4] = 45

	var i = 0
	for i < (len(array)) {
		fmt.Printf("Value is %v\n", array[i])

		if i < len(jack) {
			fmt.Printf("Value is %v\n", jack[i])
		}
		i++
	}
}
