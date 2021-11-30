// 面试题 05.02. 二进制数转字符串
//二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。如果该数字无法精确地用32位以内的二进制表示，则打印“ERROR”。
//
//示例1:
//
// 输入：0.625
// 输出："0.101"
//示例2:
//
// 输入：0.1
// 输出："ERROR"
// 提示：0.1无法被二进制准确表示
//提示：
//
//32位包括输出中的"0."这两位。
//通过次数8,126提交次数11,719
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/30/21 12:24 PM
package bianry_number_to_string_lcci

import "fmt"

func printBin(num float64) string {
	multiple := num * float64(1<<30)
	if multiple-float64(int(multiple)) != 0 {
		return "ERROR"
	}
	ans := "0." + fmt.Sprintf("%.30b", int(multiple))
	for ans[len(ans)-1] == '0' {
		ans = ans[:len(ans)-1]
	}
	return ans
}

func printBin2(num float64) string {
	ans := "0."
	for num > 0 {
		if len(ans) == 32 {
			return "ERROR"
		}
		num *= 2
		if num >= 1 {
			ans += "1"
			num -= 1
		} else {
			ans += "0"
		}
	}
	return ans
}
