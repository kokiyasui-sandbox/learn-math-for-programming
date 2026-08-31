package main

import "fmt"

func main() {
	nums := []int{0, 2, 4, 5, 6, 8, 11, 12, 13, 17, 23, 25, 34, 35, 38, 42, 66, 72, 88, 99}

	idx := binary(nums, 2)
	if idx >= 0 {
		fmt.Println(idx)
	} else {
		fmt.Println("not found")
	}
}

func binary(nums []int, n int) int {
	low := 0
	high := len(nums) - 1

	var mid int

	for low <= high {
		mid = (high + low) / 2

		if nums[mid] == n {
			return mid
		} else {
			if nums[mid] > n {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
