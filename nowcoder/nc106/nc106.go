// NC106 三个数的最大乘积
// 算法知识视频讲解
//简单  通过率：35.27%  时间限制：1秒  空间限制：64M
//知识点
//数组
//数学
//题目
//题解(15)
//讨论(32)
//排行
//描述
//给定一个长度为  的无序数组  ，包含正数、负数和 0 ，请从中找出 3 个数，使得乘积最大，返回这个乘积。
//
//要求时间复杂度：  ，空间复杂度：  。
//
//数据范围：
//3 \le n \le 2 * 10^53≤n≤2∗10
//5
//
//-10^6 \le A[i] \le 10^6−10
//6
// ≤A[i]≤10
//6
//
//示例1
//输入：
//[3,4,1,2]
//复制
//返回值：
//24
//复制
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/16/21 11:46 AM
package nc106

import "math"

/**
 * 最大乘积
 * @param A int整型一维数组
 * @return long长整型
 */
func solve(A []int) int64 {
	min1, min2, max1, max2, max3 := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32, math.MinInt32
	for _, i := range A {
		if i > max1 {
			max3 = max2
			max2 = max1
			max1 = i
		} else if i > max2 {
			max3 = max2
			max2 = i
		} else if i > max3 {
			max3 = i
		}

		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
		}
	}

	p1 := int64(min1) * int64(min2) * int64(max1)
	p2 := int64(max1) * int64(max2) * int64(max3)
	if p1 > p2 {
		return p1
	}
	return p2
}
