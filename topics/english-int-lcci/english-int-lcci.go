// 面试题 16.08. 整数的英语表示
//给定一个整数，打印该整数的英文描述。
//
//示例 1:
//
//输入: 123
//输出: "One Hundred Twenty Three"
//示例 2:
//
//输入: 12345
//输出: "Twelve Thousand Three Hundred Forty Five"
//示例 3:
//
//输入: 1234567
//输出: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
//示例 4:
//
//输入: 1234567891
//输出: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"
//注意：本题与 273 题相同：https://leetcode-cn.com/problems/integer-to-english-words/
//
//通过次数2,298提交次数6,155
//
// @author xuejunc deerhunter0837@gmail.com
// @create 10/14/21 10:06 AM
package english_int_lcci

var lessThan20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var units = []string{" Billion", " Million", " Thousand", ""}
var tenUnit = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	ans := ""
	base := 1_000_000_000
	for i := 0; i < 4; i++ {
		threeBit := num / base
		if threeBit > 0 {
			if len(ans) > 0 {
				ans += " "
			}
			ans += threeBitToWords(threeBit) + units[i]
		}
		num %= base
		base /= 1000
	}
	return ans
}

// 把一个三位数转为英文
func threeBitToWords(num int) string {
	ans := ""
	hundreds := num / 100
	if hundreds > 0 {
		ans += lessThan20[hundreds] + " Hundred"
	}
	num %= 100
	tens := num / 10
	if tens >= 2 {
		if len(ans) > 0 {
			ans += " "
		}
		ans += tenUnit[tens]
		num %= 10
	}
	if num > 0 {
		if len(ans) > 0 {
			ans += " "
		}
		ans += lessThan20[num]
	}
	return ans
}
