package algorithm

func mergeSort(nums []int) []int{
	if len(nums) < 2 {
		return nums
	}
	i := len(nums)/2
	left := mergeSort(nums[0:i])
	right := mergeSort(nums[i:])
	return merge(left, right)
}

func merge(left, right []int) []int{
	result := make([]int, 0)
	m,n:=0,0
	l,r:=len(left),len(right)
	for m<l && n<r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)
	return result
}