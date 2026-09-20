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

	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	p := partition(arr, low, high)

	quickSort(arr, low, p)
	quickSort(arr, p+1, high)
}

func partition(arr []int, low int, high int) int {
	medianOfThree(arr, low, high)
	pivot := arr[low]

	i := low - 1
	j := high + 1
	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			break
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return j
}

func medianOfThree(arr []int, low, high int) {
	mid := (high + low) / 2

	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	arr[low], arr[mid] = arr[mid], arr[low]
}
