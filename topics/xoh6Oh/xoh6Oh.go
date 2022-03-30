// 剑指 Offer II 001. 整数除法
//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
//
//
//
//注意：
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1
//
//
//示例 1：
//
//输入：a = 15, b = 2
//输出：7
//解释：15/2 = truncate(7.5) = 7
//示例 2：
//
//输入：a = 7, b = -3
//输出：-2
//解释：7/-3 = truncate(-2.33333..) = -2
//示例 3：
//
//输入：a = 0, b = 1
//输出：0
//示例 4：
//
//输入：a = 1, b = 1
//输出：1
//
//
//提示:
//
//-231 <= a, b <= 231 - 1
//b != 0
//
//
//注意：本题与主站 29 题相同：https://leetcode-cn.com/problems/divide-two-integers/
//
//
//
//通过次数28,485提交次数137,152
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/30 下午2:18
package xoh6Oh

func divide(a int, b int) int {
	if a == -(1<<31) && b == -1 {
		return 1<<31 - 1
	}
	if a < 0 {
		return -divide(-a, b)
	}
	if b < 0 {
		return -divide(a, -b)
	}
	if a < b {
		return 0
	}

	sum := b
	ans := 1
	for true {
		sum += sum
		if sum > a {
			break
		} else {
			ans <<= 1
		}
	}
	return ans + divide(a-sum>>1, b)
}

// 与方法1的区别仅仅在于将递归改为了循环
func divide2(a int, b int) int {
	// 这是处理32位溢出的问题
	if a == -(1<<31) && b == -1 {
		return 1<<31 - 1
	}
	if a < 0 {
		return -divide(-a, b)
	}
	if b < 0 {
		return -divide(a, -b)
	}
	ans := 0
	sum, quotient := b, 1
	for a >= b {
		if sum<<1 > a {
			ans += quotient
			a = a - sum
			sum, quotient = b, 1
		} else {
			sum <<= 1
			quotient <<= 1
		}
	}
	return ans
}
