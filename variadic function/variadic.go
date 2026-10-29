package main

import "fmt"

func sum(nums ...int) int {
	// func sum(nums ...interface{}){
	total := 0

	for _, nums := range nums {
		total = total + nums
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}
