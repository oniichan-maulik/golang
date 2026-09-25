package main

import "fmt"

// copied by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference functyion to be passed
func changeNum(num *int) {
	*num = 5
	fmt.Println("in vchange num", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)

	// fmt.Println("memory address", &num)
	fmt.Println("after change", num)
}
