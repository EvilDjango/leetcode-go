package topic0263

/*
263. 丑数
给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 6
输出：true
解释：6 = 2 × 3
示例 2：

输入：n = 8
输出：true
解释：8 = 2 × 2 × 2
示例 3：

输入：n = 14
输出：false
解释：14 不是丑数，因为它包含了另外一个质因数 7 。
示例 4：

输入：n = 1
输出：true
解释：1 通常被视为丑数。


提示：

-231 <= n <= 231 - 1
通过次数100,370提交次数194,868

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/28/21 6:42 PM
*/

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	n = RemoveFactor(n, 2)
	n = RemoveFactor(n, 3)
	n = RemoveFactor(n, 5)
	return n == 1
}

func RemoveFactor(n, factor int) int {
	for n%factor == 0 {
		n /= factor
	}
	return n
}
