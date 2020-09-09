package main

//给出一个区间的集合，请合并所有重叠的区间。 
//
// 
//
// 示例 1: 
//
// 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 
//
// 示例 2: 
//
// 输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。 
//
// 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。 
//
// 
//
// 提示： 
//
// 
// intervals[i][0] <= intervals[i][1] 
// 
// Related Topics 排序 数组 
// 👍 596 👎 0
import (
	"fmt"
	"sort"
)

func main(){
    r := merge([][]int{{1,4},{4,5}})
    fmt.Printf("%v", r)
}
//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for _,v := range intervals{
		if len(result)==0 || result[len(result)-1][1] < v[0] {
			result = append(result, v)
		}else {
			max := result[len(result)-1][1]
			if max < v[1] {
				max = v[1]
			}
			result[len(result)-1][1] = max
		}
	}
	return result
}


//leetcode submit region end(Prohibit modification and deletion)

