// 面试题 08.05. 递归乘法
//递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。
//
//示例1:
//
// 输入：A = 1, B = 10
// 输出：10
//示例2:
//
// 输入：A = 3, B = 4
// 输出：12
//提示:
//
//保证乘法范围不会溢出
//通过次数22,981提交次数33,911
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/26/21 6:56 PM
package recursive_mulitply_lcci

import "math/bits"

func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}
	return A<<bits.OnesCount(uint(B&-B-1)) + multiply(A, B&(B-1))
}
