package leetcode_go

/*
273. 整数转换英文表示
将非负整数 num 转换为其对应的英文表示。



示例 1：

输入：num = 123
输出："One Hundred Twenty Three"
示例 2：

输入：num = 12345
输出："Twelve Thousand Three Hundred Forty Five"
示例 3：

输入：num = 1234567
输出："One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
示例 4：

输入：num = 1234567891
输出："One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"


提示：

0 <= num <= 231 - 1
通过次数11,522提交次数37,322

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/6/21 2:18 PM
*/
var Units = [...]string{"", " Thousand", " Million", " Billion"}
var Bases = [...]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
	"Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var Tens = [...]string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	ans := ""
	i := 0
	for num > 0 {
		word := threeDigitToWords(num % 1000)
		if len(word) > 0 {
			if len(ans) > 0 {
				ans = word + Units[i] + " " + ans
			} else {
				ans = word + Units[i]
			}
		}
		num /= 1000
		i++
	}
	return ans
}

func threeDigitToWords(i int) string {
	if i <= 19 {
		return Bases[i]
	}
	if i%100 == 0 {
		return Bases[i/100] + " Hundred"
	}
	ans := ""
	if i/100 > 0 {
		ans = Bases[i/100] + " Hundred "
		i %= 100
	}
	if i < 20 {
		return ans + Bases[i]
	}

	if i%10 == 0 {
		return ans + Tens[i/10-2]
	}
	if i/10 > 0 {
		ans += Tens[i/10-2] + " "
		i %= 10
	}
	return ans + Bases[i%10]
}
