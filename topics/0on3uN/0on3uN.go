// 剑指 Offer II 087. 复原 IP
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//
//
//
//示例 1：
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//示例 2：
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//示例 3：
//
//输入：s = "1111"
//输出：["1.1.1.1"]
//示例 4：
//
//输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
//示例 5：
//
//输入：s = "10203040"
//输出：["10.20.30.40","102.0.30.40","10.203.0.40"]
//
//
//提示：
//
//0 <= s.length <= 3000
//s 仅由数字组成
//
//
//注意：本题与主站 93 题相同：https://leetcode-cn.com/problems/restore-ip-addresses/
//
//通过次数3,907提交次数6,177
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/23/21 9:21 PM
package _on3uN

import "strconv"

func restoreIpAddresses(s string) []string {
	var ans []string
	var dfs func(int, int, string)
	dfs = func(index, numbers int, ip string) {
		if index == len(s) || numbers == 0 {
			if index == len(s) && numbers == 0 {
				ans = append(ans, ip)
			}
			return
		}
		// 每位数字最多3位最少一位
		if len(s)-index > 3*numbers || len(s)-index < numbers {
			return
		}

		if index > 0 {
			ip += "."
		}
		dfs(index+1, numbers-1, ip+s[index:index+1])
		if s[index] == '0' {
			return
		}
		if index+1 < len(s) {
			dfs(index+2, numbers-1, ip+s[index:index+2])
		}
		if index+2 < len(s) {
			num, _ := strconv.Atoi(s[index : index+3])
			if num <= 255 {
				dfs(index+3, numbers-1, ip+s[index:index+3])
			}
		}
	}
	dfs(0, 4, "")
	return ans
}
