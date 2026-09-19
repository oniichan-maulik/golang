package main

import "fmt"

const name = "name"
const (
	port      = 5000
	localhost = 88888
)

func main() {
	// const name string = "name"
	fmt.Println("hello")

	fmt.Println(name)
	// const (
	// 	port      = 5000
	// 	localhost = 88888
	// )

	fmt.Println(port, localhost)
}
