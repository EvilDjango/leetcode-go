// 面试题 08.08. 有重复字符串的排列组合
//有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
//
//示例1:
//
// 输入：S = "qqe"
// 输出：["eqq","qeq","qqe"]
//示例2:
//
// 输入：S = "ab"
// 输出：["ab", "ba"]
//提示:
//
//字符都是英文字母。
//字符串长度在[1, 9]之间。
//通过次数17,164提交次数23,885
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/25/21 2:28 PM
package pile_box_lcci

import (
	"sort"
)

func permutation(S string) []string {
	bytes := []byte(S)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	var res []string
	backtrack(&res, bytes, make([]bool, len(S)), "", 0)
	return res
}

func backtrack(res *[]string, bytes []byte, used []bool, curr string, index int) {
	if index == len(bytes) {
		*res = append(*res, curr)
		return
	}

	prev := -1
	for i := range bytes {
		if used[i] || prev != -1 && bytes[prev] == bytes[i] {
			continue
		}
		used[i] = true

		backtrack(res, bytes, used, curr+string(bytes[i]), index+1)
		used[i] = false
		prev = i
	}
}
