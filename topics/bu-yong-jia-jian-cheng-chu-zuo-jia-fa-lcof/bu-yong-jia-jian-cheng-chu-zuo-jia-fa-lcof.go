// 剑指 Offer 65. 不用加减乘除做加法
//写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
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
//通过次数99,119提交次数167,446
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2022/4/6 下午12:57
package bu_yong_jia_jian_cheng_chu_zuo_jia_fa_lcof

func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
