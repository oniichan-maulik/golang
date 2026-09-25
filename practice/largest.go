package main

import "fmt"

// 50, 60,49,22,4,11

func largerNum(number []int) []int {
	result := []int{}
	for _, num2 := range number {
		if num2 > 4 {
			result = append(result, num2)
		}
	}
	return result
}

func main() {
	result := []int{50, 60, 49, 22, 4, 11}
	num := largerNum(result)
	fmt.Println(num)
}
