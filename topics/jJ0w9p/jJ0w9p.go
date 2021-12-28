// 剑指 Offer II 072. 求平方根
//给定一个非负整数 x ，计算并返回 x 的平方根，即实现 int sqrt(int x) 函数。
//
//正数的平方根有两个，只输出其中的正数平方根。
//
//如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。
//
//
//
//示例 1:
//
//输入: x = 4
//输出: 2
//示例 2:
//
//输入: x = 8
//输出: 2
//解释: 8 的平方根是 2.82842...，由于小数部分将被舍去，所以返回 2
//
//
//提示:
//
//0 <= x <= 231 - 1
//
//
//注意：本题与主站 69 题相同： https://leetcode-cn.com/problems/sqrtx/
//
//通过次数6,875提交次数16,362
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/28/21 4:41 PM
package jJ0w9p

func mySqrt(x int) int {
	low, high := 0, 1<<16
	var ans int
	for low <= high {
		mid := (high-low)/2 + low
		power := mid * mid
		if power < x {
			ans = mid
			low = mid + 1
		} else if power > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	return ans
}

// 牛顿迭代
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	c := float64(x)
	x0, x1 := float64(x), float64(x)
	for {
		x1 = (x0 + c/x0) / 2
		if x0-x1 < 10e-7 {
			break
		}
		x0 = x1
	}
	return int(x0)
}
