// 面试题 08.07. 无重复字符串的排列组合
//无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
//
//示例1:
//
// 输入：S = "qwe"
// 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//示例2:
//
// 输入：S = "ab"
// 输出：["ab", "ba"]
//提示:
//
//字符都是英文字母。
//字符串长度在[1, 9]之间。
//通过次数22,360提交次数27,476
//
// @author xuejunc deerhunter0837@gmail.com
// @create 11/25/21 4:17 PM
package permutation_i_lcci

// 回溯
func permutation(S string) []string {
	var res []string
	var backtrack func(options []byte, bytes []byte)
	backtrack = func(options []byte, bytes []byte) {
		if len(options) == 0 {
			res = append(res, string(bytes))
			return
		}
		for i, b := range options {
			bytes = append(bytes, b)
			newOptions := make([]byte, len(options))
			copy(newOptions, options)
			newOptions = append(newOptions[:i], newOptions[i+1:]...)
			backtrack(newOptions, bytes)
			bytes = bytes[:len(bytes)-1]
		}
	}
	backtrack([]byte(S), make([]byte, 0, len(S)))
	return res
}

func permutation2(S string) []string {
	var res []string
	var backtrack func(bytes []byte, used []bool, curr string)
	backtrack = func(bytes []byte, used []bool, curr string) {
		if len(curr) == len(bytes) {
			res = append(res, curr)
			return
		}
		for i := range bytes {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(bytes, used, curr+string(bytes[i]))
			used[i] = false
		}
	}
	backtrack([]byte(S), make([]bool, len(S)), "")
	return res
}
