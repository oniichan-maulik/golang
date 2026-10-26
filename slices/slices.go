package main

import (
	"fmt"
)

func main() {
	// uninitialised slice is nill
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// capacity is maximum numbers of elements to be fit
	// var nums = make([]int, 0, 5)
	// fmt.Println(cap(nums))
	// nums := []int{}
	// nums = append(nums, 1, 2)
	// nums = append(nums, 4)
	// nums = append(nums, 3)
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// copy funnction
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator
	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[0:])
	// fmt.Println(nums[:1])

	// slice package
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2))

	// 2d slice
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
}
