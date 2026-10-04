// Command h3load is the HTTP/3 load tester CLI (skeleton for Stage 0).
//
// Executables in Go live in package "main". Only this package can define
// func main(), which the runtime calls when the program starts.
package main

import "fmt"

// version is a compile-time constant we will print at startup.
// Stage 1 will revisit how we expose this (logging, -version flag, etc.).
const version = "0.0.0-dev"

func main() {

	fmt.Printf("h3load %s", version)
	
}
