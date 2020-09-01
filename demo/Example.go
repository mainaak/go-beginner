package main

//go run "complete_path" to run .exe
//go build "package path" to build file containing main
//go install "package_path" to create an executable in bin folder
//Uppercase first letter variables are exported unless in a block
//Lowercase first letter variables are package level
//Use small variable names for variables with short lifespan
//Use explanatory names for variables which have a lifespan
//Use acronyms as they are and do not camel case them forcibly

var outScope int = 100

func main() {
	channelFunction()
}
