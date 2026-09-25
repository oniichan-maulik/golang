package main

import "fmt"

// func add(a, b int) int {
// 	num1 := a
// 	num2 := b
// 	sum := num1 + num2
// 	return sum
// }

func sortSlice(demo []int) []int {
	new2 := []int{demo[1]}
	for i, new2 := range demo[0] {
		if new1 < demo[i+1] {
			new2 = append(new2, new1)
		}
	}
	return new2
}

func main() {
	// fmt.Println(add(2, 3))
	new := []int{5, 8, 3, 7, 9, 1}
	result := sortSlice(new)
	fmt.Println(result)

}
