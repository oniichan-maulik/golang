package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// sum := 0

	// i - index
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// 	// sum = sum + num
	// }
	// fmt.Println(sum)
	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// unicode code point rune
	// starting point of rune
	// 255 -> 1 byte, 255+ -> 2 byte
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
