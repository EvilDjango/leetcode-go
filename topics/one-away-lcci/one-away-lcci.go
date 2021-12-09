// 面试题 01.05. 一次编辑
//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
//示例 1:
//
//输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//示例 2:
//
//输入:
//first = "pales"
//second = "pal"
//输出: False
//通过次数35,032提交次数106,309
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/9/21 7:46 PM
package one_away_lcci

func oneEditAway(first string, second string) bool {
	return first == second || equalsOneOnReplacement(first, second) || equalsOnOneDeletion(first, second) || equalsOnOneDeletion(second, first)
}

func equalsOnOneDeletion(first string, second string) bool {
	if len(first) != len(second)+1 {
		return false
	}
	diff := 0
	for i, j := 0, 0; j < len(second); i, j = i+1, j+1 {
		if first[i] != second[j] {
			diff++
			j--
			if diff > 1 {
				return false
			}
		}
	}
	return true
}

func equalsOneOnReplacement(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}
	diff := 0
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			diff++
			if diff > 1 {
				break
			}
		}
	}
	return diff == 1
}

func oneEditAway2(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		m, n = n, m
		first, second = second, first
	}
	if m-n > 1 {
		return false
	}
	diff := 0
	for i, j := 0, 0; j < n; i, j = i+1, j+1 {
		if first[i] != second[j] {
			diff++
			if diff > 1 {
				return false
			}
			// 长度不等，那么只可能是较长的删除一个元素，所以j减一对齐
			if m != n {
				j--
			}
		}
	}
	return diff < 2
}
