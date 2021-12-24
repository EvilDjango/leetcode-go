// 剑指 Offer II 080. 含有 k 个元素的组合
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//
//
//
//示例 1:
//
//输入: n = 4, k = 2
//输出:
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
//示例 2:
//
//输入: n = 1, k = 1
//输出: [[1]]
//
//
//提示:
//
//1 <= n <= 20
//1 <= k <= n
//
//
//注意：本题与主站 77 题相同： https://leetcode-cn.com/problems/combinations/
//
//通过次数4,430提交次数5,383
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 2:57 PM
package uUsW3B

func combine(n int, k int) [][]int {
	var ans [][]int
	var dfs func(int, int, []int)
	dfs = func(i, rest int, nums []int) {
		if rest == 0 {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for j := i; j+rest-1 <= n; j++ {
			dfs(j+1, rest-1, append(nums, j))
		}
	}
	dfs(1, k, nil)
	return ans
}
