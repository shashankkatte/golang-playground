package main
// Package is like a project  or workspace
// Every file needs to fall into a package
// main word means its executable package

// Gives this package access to all the code in other package like fmt. fmt is a standard library - format
// its like stdio in c/c++
import "fmt"

// Every executable package needs this main function its the entry point to executable
func main() {
	fmt.Println("Hello, Gophers! How's it going?!")
}

// Types of packages
// 1. Executables
// 2. Reusable

// To execute this file
// 1. open a terminal in this folder
// 2. Type go run main.go
// 3. You should see the output - Hello, Gophers! How's it going?!

// Congratulations! you just ran your first go application.
