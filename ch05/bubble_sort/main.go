package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("input sequence")
		os.Exit(1)
	}

	nums := make([]int, 0, len(args))
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("inputs must be integer")
		}
		nums = append(nums, num)
	}

	nums = bubbleSort(nums)

	fmt.Println(nums)
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}

	return nums
}
