package main

import (
	"fmt"
	"strconv"
)

//1,3,5,8,8,9
func searchFirst(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			for mid > 0 {
				if nums[mid-1] != nums[mid] {
					return mid
				}
				mid--
			}
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
	return 0
}

func main() {
	result := searchFirst([]int{1, 3, 5, 8, 8, 8, 9}, 8)
	fmt.Printf(strconv.Itoa(result))
}
