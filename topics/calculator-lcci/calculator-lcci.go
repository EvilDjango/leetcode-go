// 面试题 16.26. 计算器
//给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
//
//表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
//
//示例 1:
//
//输入: "3+2*2"
//输出: 7
//示例 2:
//
//输入: " 3/2 "
//输出: 1
//示例 3:
//
//输入: " 3+5 / 2 "
//输出: 5
//说明：
//
//你可以假设所给定的表达式都是有效的。
//请不要使用内置的库函数 eval。
//通过次数14,284提交次数36,805
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/22/21 11:53 AM
package calculator_lcci

import "strconv"

func calculate(s string) int {
	ans, last, start, end := 0, 0, -1, -1
	symbol := byte('+')
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			if end == -1 {
				end = i
			}
			num, _ := strconv.Atoi(s[start:end])
			switch symbol {
			case '+':
				ans += last
				last = num
			case '-':
				ans += last
				last = -num
			case '*':
				last *= num
			case '/':
				last /= num
			}
			start, end = -1, -1
			if i != len(s) {
				symbol = s[i]
			}
		} else if s[i] != ' ' && start == -1 {
			start = i
		} else if s[i] == ' ' && start != -1 && end == -1 {
			end = i
		}
	}
	return ans + last
}
