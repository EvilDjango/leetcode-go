// 剑指 Offer II 098. 路径的数目
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
//
//
//示例 1：
//
//
//
//输入：m = 3, n = 7
//输出：28
//示例 2：
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//示例 3：
//
//输入：m = 7, n = 3
//输出：28
//示例 4：
//
//输入：m = 3, n = 3
//输出：6
//
//
//提示：
//
//1 <= m, n <= 100
//题目数据保证答案小于等于 2 * 109
//
//
//注意：本题与主站 62 题相同： https://leetcode-cn.com/problems/unique-paths/
//
//通过次数3,441提交次数4,576
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/19/21 1:37 PM
package _AoeFn

// 排列组合问题，总共需要走m+n-2步，答案是 m+n-2中选取m-1的组合数
func uniquePaths(m int, n int) int {
	if m < n {
		m, n = n, m
	}
	ans := 1
	for i := m; i <= m+n-2; i++ {
		ans *= i
	}
	for i := 2; i <= n-1; i++ {
		ans /= i
	}
	return ans
}

func Factorial(i int) int {
	ans := 1
	for j := 2; j <= i; j++ {
		ans *= j
	}
	return ans
}
