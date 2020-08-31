package main

import "fmt"

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

//todo
func selectSort(nums []int)  {

}

func main() {
	nums := []int{4,5,6,3,2,1}
	insertSort(nums)
	fmt.Print(nums)
}