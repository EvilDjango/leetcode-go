// 剑指 Offer II 002. 二进制加法
//给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//
//输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
//示例 1:
//
//输入: a = "11", b = "10"
//输出: "101"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//提示：
//
//每个字符串仅由字符 '0' 或 '1' 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 "0" ，就都不含前导零。
//
//
//注意：本题与主站 67 题相同：https://leetcode-cn.com/problems/add-binary/
//
//通过次数24,909提交次数44,732
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/3/30 下午1:35
package JFETK5

import "strconv"

// 模拟从低到的加法运算
func addBinary(a string, b string) string {
	m, n := len(a)-1, len(b)-1
	add := false
	ans := ""
	for add || m >= 0 || n >= 0 {
		sum := 0
		if m >= 0 && a[m] == '1' {
			sum += 1
		}
		if n >= 0 && b[n] == '1' {
			sum += 1
		}
		if add {
			sum += 1
		}
		add = sum > 1
		sum %= 2
		ans = strconv.Itoa(sum) + ans
		m--
		n--
	}
	return ans
}

//将a和b转换为整形数字，然后使用位运算求和
// 遇到超长的字符串会因为溢出而出错。未AC
func addBinary2(a string, b string) string {
	x, _ := strconv.ParseInt(a, 2, 0)
	y, _ := strconv.ParseInt(b, 2, 0)
	for y != 0 {
		x, y = x^y, (x&y)<<1
	}
	return strconv.FormatInt(x, 2)
}
