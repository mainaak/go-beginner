package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type todos struct {
	userId int
	id int
	title string
	completed bool
}

//Defer will execute the statement next to it after the block, it runs in reverse and maintains the state at the point at which it was triggered
//Defers  execute after panic, hence ensuring that the open resources are closed before the panic
func makeACall() {

	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if error != nil {
		log.Fatal(error)
	}

	json, error := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Response: %v", string(json))
}
