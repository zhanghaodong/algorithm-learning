package src

func quickSort(nums []int)  {
	quickSortHelper(nums,0, len(nums)-1)
}

func quickSortHelper(nums []int, start, end int)  {
	if start >= end {
		return
	}
	position := partition(nums, start, end)
	quickSortHelper(nums, start, position-1)
	quickSortHelper(nums, position+1, end)
}

/*
	6,11,3,9,8
	6 < 8 swap
	11 > 8 no swap
	3 < 8 swap
	9 > 8 no swap
*/
func partition(nums []int, start, end int)  int{
	pivot := nums[end]
	left := start
	for i:=start; i< end; i++ {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[end], nums[left] = nums[left], nums[end]
	return left
}