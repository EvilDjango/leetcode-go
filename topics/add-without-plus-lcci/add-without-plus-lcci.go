// 面试题 17.01. 不用加号的加法
//设计一个函数把两个数字相加。不得使用 + 或者其他算术运算符。
//
//示例:
//
//输入: a = 1, b = 1
//输出: 2
//
//
//提示：
//
//a, b 均可能是负数或 0
//结果不会溢出 32 位整数
//通过次数12,033提交次数20,153
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/21/21 2:44 PM
package add_without_plus_lcci

func add(a int, b int) int {
	a, b = a^b, a&b<<1
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
