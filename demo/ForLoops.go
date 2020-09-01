package main

import "fmt"

func execForLoop(countOfLoop int) {

	for q := 0; q < countOfLoop; q++ {
		fmt.Printf("Loop Number %v\n", q+1)
	}

}

func continueInLoops() {

	for q := 0; q < 5; q++ {

		if q%2 != 0 {
			fmt.Printf("Iteration %v will break flow\n", q)
			continue
		}

		fmt.Printf("Loop Iteration Unskipped For %v\n", q)
	}
}

func coolLoops() {

JuliusCaesar:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (i * j) >= 10 {
				break JuliusCaesar
			}
			fmt.Printf("Printing for [%v][%v]\n", i, j)
		}
	}
}

func generateEvenNumbers(length int) {
	tableArray := [20]int{}
	for i, j := 0, 0; j < length; i++ {
		if i%2 == 0 && i != 0 {
			tableArray[j] = i
			j++
		}
	}

	//We can use _ instead of k if we only want to use the value and not the key
	for k, v := range tableArray {
		if k >= length {
			break
		}
		fmt.Printf("2 x %v = %v\n", k+1, v)
	}
}
