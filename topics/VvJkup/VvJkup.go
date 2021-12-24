// 剑指 Offer II 083. 没有重复元素集合的全排列
//给定一个不含重复数字的整数数组 nums ，返回其 所有可能的全排列 。可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
//
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同
//
//
//注意：本题与主站 46 题相同：https://leetcode-cn.com/problems/permutations/
//
//通过次数5,709提交次数6,631
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/24/21 12:29 AM
package VvJkup

func permute(nums []int) [][]int {
	n := len(nums)
	visited := make([]bool, n)
	var ans [][]int
	var dfs func(int, []int)
	dfs = func(count int, path []int) {
		if count == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				visited[i] = true
				dfs(count+1, append(path, nums[i]))
				visited[i] = false
			}
		}
	}
	dfs(0, nil)
	return ans
}
