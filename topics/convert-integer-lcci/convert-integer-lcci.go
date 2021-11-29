// 面试题 05.06. 整数转换
//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
//
//示例1:
//
// 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
// 输出：2
//示例2:
//
// 输入：A = 1，B = 2
// 输出：2
//提示:
//
//A，B范围在[-2147483648, 2147483647]之间
//通过次数12,275提交次数23,205
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/29/21 9:10 PM
package convert_integer_lcci

func convertInteger(A int, B int) int {
	c := int32(A) ^ int32(B)
	count := 0
	for c != 0 {
		count++
		c &= c - 1
	}
	return count
}
