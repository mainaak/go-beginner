package main

import (
	"fmt"
	"strconv"
)

func maps() {
	applicantDetails := make(map[string]string)
	applicantDetails = map[string]string{
		"applicantId":    strconv.Itoa(007),
		"clientId":       strconv.Itoa(520),
		"applicantName":  "Mainaak Aanand",
		"applicantEmail": "mainaak@gmail.com",
	}
	applicantDetails["addedKey"] = "Testing Value"
	applicantDetails["applicantEmail"] = "Sharma"
	delete(applicantDetails, "clientId")
	fmt.Printf("The map is\n%v\n", applicantDetails)

	_, ok := applicantDetails["clientId"]

	if !ok {
		fmt.Printf("The key clientId does not exist\n")
	} else {
		fmt.Printf("The value for the key clientId is %v", applicantDetails["clientId"])
	}
}
