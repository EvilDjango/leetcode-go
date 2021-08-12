package topic282

import "strconv"

/*
282. 给表达式添加运算符
给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。



示例 1:

输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]
示例 2:

输入: num = "232", target = 8
输出: ["2*3+2", "2+3*2"]
示例 3:

输入: num = "105", target = 5
输出: ["1*0+5","10-5"]
示例 4:

输入: num = "00", target = 0
输出: ["0+0", "0-0", "0*0"]
示例 5:

输入: num = "3456237490", target = 9191
输出: []


提示：

0 <= num.length <= 10
num 仅含数字
通过次数6,770提交次数17,476

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/12/21 9:00 PM
*/
func addOperators(num string, target int) []string {
	var ans []string
	dfs(num, target, &ans, 0, 0, 0, "")
	return ans
}

func dfs(num string, target int, ans *[]string, index, sum, last int, s string) {
	if index == len(num) {
		if sum+last == target {
			*ans = append(*ans, s)
		}
		return
	}

	for i := index; i < len(num); i++ {
		sub := num[index : i+1]
		d, _ := strconv.Atoi(sub)
		if index == 0 {
			dfs(num, target, ans, i+1, 0, d, sub)
		} else {
			dfs(num, target, ans, i+1, sum+last, d, s+"+"+sub)
			dfs(num, target, ans, i+1, sum+last, -d, s+"-"+sub)
			dfs(num, target, ans, i+1, sum, last*d, s+"*"+sub)
		}
		if num[index] == '0' {
			break
		}
	}
}
