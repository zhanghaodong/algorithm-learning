package main

import "sort"

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针 
// 👍 2564 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int,0)

	for first:=0; first<n;first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		if nums[first] > 0{
			break
		}
		third := n-1
		target := -1*nums[first]
		for second:=first+1; second<n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[third] + nums[second] > target {
				third--
			}
			if second == third{
				break
			}
			if nums[second] + nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)

