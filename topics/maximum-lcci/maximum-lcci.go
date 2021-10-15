// 面试题 16.07. 最大数值
//编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。
//
//示例：
//
//输入： a = 1, b = 2
//输出： 2
//通过次数19,663提交次数26,825
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/15/21 7:29 AM
package maximum_lcci

func maximum(a int, b int) int {
	k := (a - b) >> 63
	return a*(k+1) - b*k
}
