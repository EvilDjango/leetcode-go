// 剑指 Offer 64. 求1+2+…+n
//求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
//
//
//
//示例 1：
//
//输入: n = 3
//输出: 6
//示例 2：
//
//输入: n = 9
//输出: 45
//
//
//限制：
//
//1 <= n <= 10000
//通过次数176,081提交次数205,190
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/4/6 下午1:04
package qiu_12n_lcof

// 递归+短路与
func sumNums(n int) int {
	ans := 0
	var dfs func(int) bool
	dfs = func(i int) bool {
		ans += i
		return i > 0 && dfs(i-1)
	}
	dfs(n)
	return ans
}
