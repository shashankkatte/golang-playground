// Package is like a project  or workspace
// Every file needs to fall into a package
// main word means its executable package
package main

// Gives this package access to all the code in other package like fmt. fmt is a standard library - format
// its like stdio in c/c++
import "fmt"

// Every executable package needs this main function its the entry point to executable
func main() {
	fmt.Println("Hello Go!")
}

// Types of packages
// 1. Executables
// 2. Reusable
