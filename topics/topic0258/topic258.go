package topic0258

/*
258. 各位相加
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

示例:

输入: 38
输出: 2
解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
进阶:
你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？

通过次数73,876提交次数106,986

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 10:41 AM
*/

// 暴力法
func addDigits(num int) int {
	for num >= 10 {
		num = addDigitsOnce(num)
	}
	return num
}

func addDigitsOnce(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// 数学法
func addDigits2(num int) int {
	return (num - 1%9) + 1
}
