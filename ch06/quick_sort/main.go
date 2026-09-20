package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: main <file name> \n")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot convert number: %v\n", err)
			os.Exit(1)
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	nums = quickSort(nums)
	fmt.Println(nums)
}

func quickSort(arr []int) []int {
	var high []int
	var same []int
	var low []int

	if len(arr) < 2 {
		return arr
	}

	pivot := selectPivot(arr)

	for _, value := range arr {
		if value > pivot {
			high = append(high, value)
		} else if value == pivot {
			same = append(same, value)
		} else if value < pivot {
			low = append(low, value)
		}
	}

	ret := append(quickSort(low), same...)
	ret = append(ret, quickSort(high)...)

	return ret
}

func selectPivot(arr []int) int {
	first := arr[0]
	end := arr[len(arr)-1]
	middle := arr[(len(arr)-1)/2]

	if (first <= middle && middle <= end) || (end <= middle && middle <= first) {
		return middle
	}
	if (middle <= first && first <= end) || (end <= first && first <= middle) {
		return first
	}
	return end
}
