// kmp算法
//
// @author xuejunc deerhunter0837@gmail.com
// @create 9/15/21 2:09 PM
package common

// 搜索s[l:r]中首次出现字符串sub的位置,不存在返回-1
func search(s, sub string, l, r int) int {
	if r-l < len(sub) {
		return -1
	}
	next := getNext(sub)
	i, j := l, 0
	for i < r && j < len(sub) {
		if j == -1 || s[i] == sub[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(sub) {
		return i - j
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	i, j := 0, -1
	for i < len(s) {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
