package topic254

import (
	"math"
)

/*
254. 因子的组合
整数可以被看作是其因子的乘积。

例如：

8 = 2 x 2 x 2;
  = 2 x 4.
请实现一个函数，该函数接收一个整数 n 并返回该整数所有的因子组合。

注意：

你可以假定 n 为永远为正数。
因子必须大于 1 并且小于 n。
示例 1：

输入: 1
输出: []
示例 2：

输入: 37
输出: []
示例 3：

输入: 12
输出:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
示例 4:

输入: 32
输出:
[
  [2, 16],
  [2, 2, 8],
  [2, 2, 2, 4],
  [2, 2, 2, 2, 2],
  [2, 4, 4],
  [4, 8]
]
通过次数4,470提交次数7,851

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/27/21 10:48 AM
*/

// 暴力算法

func getFactors(n int) [][]int {
	return doGetFactors(n, 2)
}

func doGetFactors(n, prev int) [][]int {
	var ans [][]int
	for i := prev; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			ans = append(ans, []int{n / i, i})
			subs := doGetFactors(n/i, i)
			for _, sub := range subs {
				sub = append(sub, i)
				ans = append(ans, sub)
			}
		}
	}
	return ans
}
