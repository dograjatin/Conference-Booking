// All the go programs must be a part of a package and the standard name for a package is main
package main

// All the functions in the go are part of some package so the print function used here is part of
// fmt i.e. format package. In go we need to explicitely import every package
import "fmt"

// Go needs an entry point to know from where to start executing and the main entry
// point is the function main
func Hello() {

	fmt.Println("Hello World")
}

// To run Go program all the files must be a part of a project so to initiate a project
// run "go mod init <project-name>"
// then to run a go file/program/script run "go run <file-name>"
