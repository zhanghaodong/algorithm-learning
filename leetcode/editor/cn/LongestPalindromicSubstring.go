package main

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。 
//
// 示例 1： 
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
// 
//
// 示例 2： 
//
// 输入: "cbbd"
//输出: "bb"
// 
// Related Topics 字符串 动态规划 
// 👍 2663 👎 0


func main(){
    println(longestPalindrome("babad"))
}
//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	n:=len(s)
	if n < 2 {
		return s
	}
	maxLen,begin := 1,0
	dp := make([][]bool, n)
	for i:=0;i<n;i++ {
		dp[i] = make([]bool, n)
	}

	for j:=1; j<n; j++ {
		for i:=0; i<n; i++ {
			if s[j] != s[i] {
				dp[i][j] = false
			}else {
				if j-i<3 {
					dp[i][j] = true
				}else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLen{
				maxLen = j-i+1
				begin = i
			}
		}

	}
	return s[begin:begin+maxLen]
}
//leetcode submit region end(Prohibit modification and deletion)

