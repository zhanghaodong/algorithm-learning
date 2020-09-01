package main

func bubbleSort(nums []int) {
	if len(nums) > 1 {
		for k:=0; k<len(nums);k++ {
			flag := false
			for i:=0; i<len(nums)-k-1; i++{
				if nums[i] > nums[i+1] {
					temp := nums[i+1]
					nums[i+1] = nums[i]
					nums[i] = temp
					flag = true
				}
			}
			if !flag {
				break
			}
		}
	}
}

func insertSort(nums []int)  {
	if len(nums) <= 1{
		return
	}
	for i:=1;i<len(nums);i++ {
		j := i-1
		value := nums[i]
		for ;j>=0;j--{
			if value < nums[j] {
				nums[j+1] = nums[j]
			}else {
				break
			}
		}
		nums[j+1] = value
	}
}

func selectSort(nums []int)  {
	if len(nums) <=1 {
		return
	}
	for i:=0;i<len(nums)-1;i++ {
		minIndex := len(nums)-1
		minValue := nums[minIndex]
		for k:=i;k<len(nums);k++ {
			if minValue > nums[k] {
				minValue = nums[k]
				minIndex = k
			}
		}
		temp := nums[minIndex]
		nums[minIndex] = nums[i]
		nums[i] = temp
	}
}
