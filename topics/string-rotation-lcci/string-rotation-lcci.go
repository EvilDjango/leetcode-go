// 面试题 01.09. 字符串轮转
//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
//
//示例1:
//
// 输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//示例2:
//
// 输入：s1 = "aa", s2 = "aba"
// 输出：False
//提示：
//
//字符串长度在[0, 100000]范围内。
//说明:
//
//你能只调用一次检查子串的方法吗？
//通过次数35,504提交次数64,744
//
// @author xuejunc deerhunter0837@gmail.com
// @create 12/8/21 10:25 PM
package string_rotation_lcci

import "strings"

// 暴力算法
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == "" && s2 == "" {
		return true
	}
	for i := 0; i < len(s1); i++ {
		if s1[i:]+s1[:i] == s2 {
			return true
		}
	}
	return false
}

func isFlipedString2(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}
