package topic248

/*
248. 中心对称数 III
中心对称数是指一个数字在旋转了 180 度之后看起来依旧相同的数字（或者上下颠倒地看）。

写一个函数来计算范围在 [low, high] 之间中心对称数的个数。

示例:

输入: low = "50", high = "100"
输出: 3
解释: 69，88 和 96 是三个在该范围内的中心对称数
注意:
由于范围可能很大，所以 low 和 high 都用字符串表示。

通过次数2,444提交次数5,241

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   7/22/21 3:37 PM
*/

func strobogrammaticInRange(low string, high string) int {
	ret := 0
	m, n := len(low), len(high)
	for i := m; i <= n; i++ {
		strs := strobogrammatic(i, i)
		if i > m && i < n {
			ret += len(strs)
			continue
		}
		for _, str := range strs {
			if compare(str, low) >= 0 && compare(high, str) >= 0 {
				ret++
			}
		}
	}
	return ret
}

func strobogrammatic(len, remain int) []string {
	if remain == 0 {
		return []string{""}
	} else if remain == 1 {
		return []string{"0", "1", "8"}
	}
	var ret []string
	subs := strobogrammatic(len, remain-2)
	for _, sub := range subs {
		if len != remain {
			ret = append(ret, "0"+sub+"0")
		}
		ret = append(ret, "1"+sub+"1")
		ret = append(ret, "6"+sub+"9")
		ret = append(ret, "8"+sub+"8")
		ret = append(ret, "9"+sub+"6")
	}
	return ret
}

// 比较两数字大小
func compare(i, j string) int {
	m, n := len(i), len(j)
	if m > n {
		return 1
	}
	if m < n {
		return -1
	}
	if i > j {
		return 1
	}
	if i < j {
		return -1
	}
	return 0
}
