// 面试题 16.05. 阶乘尾数
//设计一个算法，算出 n 阶乘有多少个尾随零。
//
//示例 1:
//
//输入: 3
//输出: 0
//解释: 3! = 6, 尾数中没有零。
//示例 2:
//
//输入: 5
//输出: 1
//解释: 5! = 120, 尾数中有 1 个零.
//说明: 你算法的时间复杂度应为 O(log n) 。
//
//通过次数10,832提交次数23,334
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/15/21 8:51 AM
package factorial_zeros_lcci

func trailingZeroes(n int) int {
	cnt := 0
	for n >= 5 {
		n /= 5
		cnt += n
	}
	return cnt
}
