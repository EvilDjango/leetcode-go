package topic0248

/*
248题第二种解法，使用队列，从底到顶。

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/22/21 6:04 PM
*/

func strobogrammaticInRange2(low string, high string) int {
	ret := 0
	queue := []string{"", "0", "1", "8"}
	for len(queue) > 0 {
		num := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// 要排除多位数以0开头的情形
		if !(len(num) > 1 && num[0] == '0') && gte(num, low) && gte(high, num) {
			ret++
		}
		if len(num)+2 <= len(high) {
			queue = append(queue, "0"+num+"0")
			queue = append(queue, "1"+num+"1")
			queue = append(queue, "6"+num+"9")
			queue = append(queue, "8"+num+"8")
			queue = append(queue, "9"+num+"6")
		}
	}
	return ret
}

// 判断num1 是否 大于等于 num2
func gte(num1, num2 string) bool {
	if len(num1) != len(num2) {
		return len(num1) > len(num2)
	}
	return num1 >= num2
}
