package main

import "fmt"

type class struct {
	Name  string
	Age   int
	Marks float32
}

func main() {
	myName := class{
		Name:  "Maulik",
		Age:   22,
		Marks: 84,
	}

	fmt.Println(myName)
}
