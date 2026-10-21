package main

import "fmt"

const outside = "Outside"

func main() {
	const name = "Yash"

	// name = "abc" // thorws error

	const (
		port = 9000
		host = "localhost"
	)

	fmt.Println(name, port, host, outside)
}
