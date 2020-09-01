package main

import (
	"net/http"
)

func callback(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("I am a reponse"))
}

//Turn on another application on port 8080 before running this
//panic is like exception in java
func makeListener(path string){
	http.HandleFunc(path, callback)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic("CANNOT TOUCH THIS\n" + err.Error())
	}
}
